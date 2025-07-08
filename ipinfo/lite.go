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
	defaultLiteBaseURL = "https://api.ipinfo.io/lite/"
)

// LiteClient is a client for the IPinfo Lite API.
type LiteClient struct {
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

// Lite represents the response from the IPinfo Lite API.
type Lite struct {
	IP            net.IP `json:"ip"`
	ASN           string `json:"asn"`
	ASName        string `json:"as_name"`
	ASDomain      string `json:"as_domain"`
	CountryCode   string `json:"country_code"`
	Country       string `json:"country"`
	ContinentCode string `json:"continent_code"`
	Continent     string `json:"continent"`
	Bogon         bool   `json:"bogon"`

	// Extended fields using the same country data as Core API
	CountryName     string          `json:"-"`
	CountryFlag     CountryFlag     `json:"-"`
	CountryFlagURL  string          `json:"-"`
	CountryCurrency CountryCurrency `json:"-"`
	ContinentInfo   Continent       `json:"-"`
	IsEU            bool            `json:"-"`
}

func (v *Lite) setCountryName() {
	if v.CountryCode != "" {
		v.CountryName = GetCountryName(v.CountryCode)
		v.IsEU = IsEU(v.CountryCode)
		v.CountryFlag.Emoji = GetCountryFlagEmoji(v.CountryCode)
		v.CountryFlag.Unicode = GetCountryFlagUnicode(v.CountryCode)
		v.CountryFlagURL = GetCountryFlagURL(v.CountryCode)
		v.CountryCurrency.Code = GetCountryCurrencyCode(v.CountryCode)
		v.CountryCurrency.Symbol = GetCountryCurrencySymbol(v.CountryCode)
		v.ContinentInfo.Code = GetContinentCode(v.CountryCode)
		v.ContinentInfo.Name = GetContinentName(v.CountryCode)
	}
}

// NewLiteClient creates a new IPinfo Lite API client.
func NewLiteClient(httpClient *http.Client, cache *Cache, token string) *LiteClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultLiteBaseURL)
	return &LiteClient{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
		Cache:     cache,
		Token:     token,
	}
}

// GetIPInfo returns the lite details for the specified IP.
func (c *LiteClient) GetIPInfo(ip net.IP) (*Lite, error) {
	if ip != nil && isBogon(netip.MustParseAddr(ip.String())) {
		bogonResponse := new(Lite)
		bogonResponse.Bogon = true
		bogonResponse.IP = ip
		return bogonResponse, nil
	}
	relUrl := "me"
	if ip != nil {
		relUrl = ip.String()
	}

	if c.Cache != nil {
		if res, err := c.Cache.Get(cacheKey(relUrl)); err == nil {
			return res.(*Lite), nil
		}
	}

	req, err := c.newRequest(nil, "GET", relUrl, nil)
	if err != nil {
		return nil, err
	}

	res := new(Lite)
	if _, err := c.do(req, res); err != nil {
		return nil, err
	}

	res.setCountryName()

	if c.Cache != nil {
		if err := c.Cache.Set(cacheKey(relUrl), res); err != nil {
			return res, err
		}
	}

	return res, nil
}

func (c *LiteClient) newRequest(ctx context.Context,
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

func (c *LiteClient) do(
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

// GetIPInfo returns the details for the specified IP.
func GetIPInfoLite(ip net.IP) (*Lite, error) {
	return DefaultLiteClient.GetIPInfo(ip)
}
