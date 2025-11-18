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
	defaultPlusBaseURL = "https://api.ipinfo.io/lookup/"
)

// PlusClient is a client for the IPinfo Plus API.
type PlusClient struct {
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

// Plus represents the response from the IPinfo Plus API /lookup endpoint.
type Plus struct {
	IP          net.IP         `json:"ip"`
	Hostname    string         `json:"hostname,omitempty"`
	Bogon       bool           `json:"bogon,omitempty"`
	Geo         *PlusGeo       `json:"geo,omitempty"`
	AS          *PlusAS        `json:"as,omitempty"`
	Mobile      *PlusMobile    `json:"mobile,omitempty"`
	Anonymous   *PlusAnonymous `json:"anonymous,omitempty"`
	IsAnonymous bool           `json:"is_anonymous"`
	IsAnycast   bool           `json:"is_anycast"`
	IsHosting   bool           `json:"is_hosting"`
	IsMobile    bool           `json:"is_mobile"`
	IsSatellite bool           `json:"is_satellite"`
	Abuse       *PlusAbuse     `json:"abuse,omitempty"`
	Company     *PlusCompany   `json:"company,omitempty"`
	Privacy     *PlusPrivacy   `json:"privacy,omitempty"`
	Domains     *PlusDomains   `json:"domains,omitempty"`
}

// PlusGeo represents the geo object in Plus API response.
type PlusGeo struct {
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
	DMACode       string  `json:"dma_code,omitempty"`
	GeonameID     string  `json:"geoname_id,omitempty"`
	Radius        int     `json:"radius"`
	LastChanged   string  `json:"last_changed,omitempty"`

	// Extended fields using the same country data as legacy Core API
	CountryName     string          `json:"-"`
	IsEU            bool            `json:"-"`
	CountryFlag     CountryFlag     `json:"-"`
	CountryFlagURL  string          `json:"-"`
	CountryCurrency CountryCurrency `json:"-"`
	ContinentInfo   Continent       `json:"-"`
}

// PlusAS represents the AS object in Plus API response.
type PlusAS struct {
	ASN         string `json:"asn"`
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	Type        string `json:"type"`
	LastChanged string `json:"last_changed,omitempty"`
}

// PlusMobile represents the mobile object in Plus API response.
type PlusMobile struct {
	Name string `json:"name,omitempty"`
	MCC  string `json:"mcc,omitempty"`
	MNC  string `json:"mnc,omitempty"`
}

// PlusAnonymous represents the anonymous object in Plus API response.
type PlusAnonymous struct {
	IsProxy bool   `json:"is_proxy"`
	IsRelay bool   `json:"is_relay"`
	IsTor   bool   `json:"is_tor"`
	IsVPN   bool   `json:"is_vpn"`
	Name    string `json:"name,omitempty"`
}

// PlusAbuse represents the abuse object in Plus API response.
type PlusAbuse struct {
	Address     string `json:"address,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryName string `json:"country_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	Network     string `json:"network,omitempty"`
	Phone       string `json:"phone,omitempty"`
}

// PlusCompany represents the company object in Plus API response.
type PlusCompany struct {
	Name   string `json:"name,omitempty"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
}

// PlusPrivacy represents the privacy object in Plus API response.
type PlusPrivacy struct {
	VPN     bool   `json:"vpn"`
	Proxy   bool   `json:"proxy"`
	Tor     bool   `json:"tor"`
	Relay   bool   `json:"relay"`
	Hosting bool   `json:"hosting"`
	Service string `json:"service,omitempty"`
}

// PlusDomains represents the domains object in Plus API response.
type PlusDomains struct {
	IP      string   `json:"ip,omitempty"`
	Total   uint64   `json:"total"`
	Domains []string `json:"domains,omitempty"`
}

func (v *Plus) enrichGeo() {
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
	if v.Abuse != nil && v.Abuse.Country != "" {
		v.Abuse.CountryName = GetCountryName(v.Abuse.Country)
	}
}

// NewPlusClient creates a new IPinfo Plus API client.
func NewPlusClient(httpClient *http.Client, cache *Cache, token string) *PlusClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultPlusBaseURL)
	return &PlusClient{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
		Cache:     cache,
		Token:     token,
	}
}

// GetIPInfo returns the Plus details for the specified IP.
func (c *PlusClient) GetIPInfo(ip net.IP) (*Plus, error) {
	if ip != nil && isBogon(netip.MustParseAddr(ip.String())) {
		bogonResponse := new(Plus)
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
			return res.(*Plus), nil
		}
	}

	req, err := c.newRequest(nil, "GET", relUrl, nil)
	if err != nil {
		return nil, err
	}

	res := new(Plus)
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

func (c *PlusClient) newRequest(ctx context.Context,
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

func (c *PlusClient) do(
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

// GetIPInfoPlus returns the Plus details for the specified IP.
func GetIPInfoPlus(ip net.IP) (*Plus, error) {
	return DefaultPlusClient.GetIPInfo(ip)
}
