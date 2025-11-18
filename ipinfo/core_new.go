package ipinfo

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/netip"
	"net/url"
	"strings"
)

const (
	defaultCoreBaseURL = "https://api.ipinfo.io/lookup/"
)

// CoreClient is a client for the IPinfo Core API.
type CoreClient struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with the IPinfo API.
	UserAgent string

	// Cache interface implementation to prevent API quota overuse.
	Cache *Cache

	// The API token used for authorization.
	Token string
}

// CoreResponse represents the response from the IPinfo Core API /lookup endpoint.
type CoreResponse struct {
	IP          net.IP   `json:"ip"`
	Bogon       bool     `json:"bogon,omitempty"`
	Geo         *CoreGeo `json:"geo,omitempty"`
	AS          *CoreAS  `json:"as,omitempty"`
	IsAnonymous bool     `json:"is_anonymous"`
	IsAnycast   bool     `json:"is_anycast"`
	IsHosting   bool     `json:"is_hosting"`
	IsMobile    bool     `json:"is_mobile"`
	IsSatellite bool     `json:"is_satellite"`
}

// CoreGeo represents the geo object in Core API response.
type CoreGeo struct {
	City          string  `json:"city,omitempty"`
	Region        string  `json:"region,omitempty"`
	RegionCode    string  `json:"region_code,omitempty"`
	Country       string  `json:"country,omitempty"`
	CountryCode   string  `json:"country_code,omitempty"`
	Continent     string  `json:"continent,omitempty"`
	ContinentCode string  `json:"continent_code,omitempty"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Timezone      string  `json:"timezone,omitempty"`
	PostalCode    string  `json:"postal_code,omitempty"`

	// Extended fields using the same country data as legacy Core API
	CountryName     string          `json:"-"`
	IsEU            bool            `json:"-"`
	CountryFlag     CountryFlag     `json:"-"`
	CountryFlagURL  string          `json:"-"`
	CountryCurrency CountryCurrency `json:"-"`
	ContinentInfo   Continent       `json:"-"`
}

// CoreAS represents the AS object in Core API response.
type CoreAS struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

func (v *CoreResponse) enrichGeo() {
	if v.Geo != nil && v.Geo.CountryCode != "" {
		v.Geo.CountryName = GetCountryName(v.Geo.CountryCode)
		v.Geo.IsEU = IsEU(v.Geo.CountryCode)
		v.Geo.CountryFlag.Emoji = GetCountryFlagEmoji(v.Geo.CountryCode)
		v.Geo.CountryFlag.Unicode = GetCountryFlagUnicode(v.Geo.CountryCode)
		v.Geo.CountryFlagURL = GetCountryFlagURL(v.Geo.CountryCode)
		v.Geo.CountryCurrency.Code = GetCountryCurrencyCode(v.Geo.CountryCode)
		v.Geo.CountryCurrency.Symbol = GetCountryCurrencySymbol(v.Geo.CountryCode)
		v.Geo.ContinentInfo.Code = GetContinentCode(v.Geo.CountryCode)
		v.Geo.ContinentInfo.Name = GetContinentName(v.Geo.CountryCode)
	}
}

// NewCoreClient creates a new IPinfo Core API client.
func NewCoreClient(httpClient *http.Client, cache *Cache, token string) *CoreClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultCoreBaseURL)
	return &CoreClient{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
		Cache:     cache,
		Token:     token,
	}
}

// GetIPInfo returns the Core details for the specified IP.
func (c *CoreClient) GetIPInfo(ip net.IP) (*CoreResponse, error) {
	if ip != nil && isBogon(netip.MustParseAddr(ip.String())) {
		bogonResponse := new(CoreResponse)
		bogonResponse.Bogon = true
		bogonResponse.IP = ip
		return bogonResponse, nil
	}
	relUrl := ""
	if ip != nil {
		relUrl = ip.String()
	}

	if c.Cache != nil {
		if res, err := c.Cache.Get(cacheKey(relUrl)); err == nil {
			return res.(*CoreResponse), nil
		}
	}

	req, err := c.newRequest(nil, "GET", relUrl, nil)
	if err != nil {
		return nil, err
	}

	res := new(CoreResponse)
	if _, err := c.do(req, res); err != nil {
		return nil, err
	}

	res.enrichGeo()

	if c.Cache != nil {
		if err := c.Cache.Set(cacheKey(relUrl), res); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (c *CoreClient) newRequest(ctx context.Context,
	method string,
	urlStr string,
	body io.Reader,
) (*http.Request, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	u := new(url.URL)
	baseURL := c.BaseURL
	if rel, err := url.Parse(urlStr); err == nil {
		u = baseURL.ResolveReference(rel)
	} else if strings.ContainsRune(urlStr, ':') {
		// IPv6 strings fail to parse as URLs, so let's add it as a URL Path.
		*u = *baseURL
		u.Path += urlStr
	} else {
		return nil, err
	}

	// get `http` package request object.
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// set common headers.
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	return req, nil
}

func (c *CoreClient) do(
	req *http.Request,
	v interface{},
) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				// ignore EOF errors caused by empty response body
				err = nil
			}
		}
	}

	return resp, err
}

// GetIPInfoCore returns the Core details for the specified IP.
func GetIPInfoCore(ip net.IP) (*CoreResponse, error) {
	return DefaultCoreClient.GetIPInfo(ip)
}
