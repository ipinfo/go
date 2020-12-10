package ipinfo

import (
	"net"
)

// Core represents data from the Core API.
type Core struct {
	IP       net.IP       `json:"ip"`
	Hostname string       `json:"hostname"`
	City     string       `json:"city"`
	Region   string       `json:"region"`
	Country  string       `json:"country"`
	Location string       `json:"loc"`
	Org      string       `json:"org"`
	Postal   string       `json:"postal"`
	Timezone string       `json:"timezone"`
	ASN      *CoreASN     `json:"asn"`
	Company  *CoreCompany `json:"company"`
	Carrier  *CoreCarrier `json:"carrier"`
	Privacy  *CorePrivacy `json:"privacy"`
	Abuse    *CoreAbuse   `json:"abuse"`
	Domains  *CoreDomains `json:"domains"`
}

// CoreASN represents ASN data for the Core API.
type CoreASN struct {
	ASN    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

// CoreCompany represents company data for the Core API.
type CoreCompany struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

// CoreCarrier represents carrier data for the Core API.
type CoreCarrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

// CorePrivacy represents privacy data for the Core API.
type CorePrivacy struct {
	Vpn     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Hosting bool `json:"hosting"`
}

// CoreAbuse represents abuse data for the Core API.
type CoreAbuse struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

// CoreDomains represents domains data for the Core API.
type CoreDomains struct {
	IP      string   `json:"ip"`
	Total   uint64   `json:"total"`
	Domains []string `json:"domains"`
}

/* CORE */

// GetIPInfo returns the details for the specified IP.
func GetIPInfo(ip net.IP) (*Core, error) {
	return DefaultClient.GetIPInfo(ip)
}

// GetIPInfo returns the details for the specified IP.
func (c *Client) GetIPInfo(ip net.IP) (*Core, error) {
	v := new(Core)
	res, err := c.getIPSpecific(ip, "json", v)
	if err != nil {
		return nil, err
	}
	return res.(*Core), nil
}

/* IP ADDRESS */

// GetIPAddr returns the IP address that IPinfo sees when you make a request.
func GetIPAddr() (string, error) {
	return DefaultClient.GetIPAddr()
}

// GetIPAddr returns the IP address that IPinfo sees when you make a request.
func (c *Client) GetIPAddr() (string, error) {
	return c.getIPSpecificStr(nil, "ip")
}

/* HOSTNAME */

// GetIPHostname returns the hostname of the domain on the specified IP.
func GetIPHostname(ip net.IP) (string, error) {
	return DefaultClient.GetIPHostname(ip)
}

// GetIPHostname returns the hostname of the domain on the specified IP.
func (c *Client) GetIPHostname(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "hostname")
}

/* CITY */

// GetIPCity returns the city for the specified IP.
func GetIPCity(ip net.IP) (string, error) {
	return DefaultClient.GetIPCity(ip)
}

// GetIPCity returns the city for the specified IP.
func (c *Client) GetIPCity(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "city")
}

/* REGION */

// GetIPRegion returns the region for the specified IP.
func GetIPRegion(ip net.IP) (string, error) {
	return DefaultClient.GetIPRegion(ip)
}

// GetIPRegion returns the region for the specified IP.
func (c *Client) GetIPRegion(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "region")
}

/* COUNTRY */

// GetIPCountry returns the country for the specified IP.
func GetIPCountry(ip net.IP) (string, error) {
	return DefaultClient.GetIPCountry(ip)
}

// GetIPCountry returns the country for the specified IP.
func (c *Client) GetIPCountry(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "country")
}

/* LOCATION */

// GetIPLocation returns the location for the specified IP.
func GetIPLocation(ip net.IP) (string, error) {
	return DefaultClient.GetIPLocation(ip)
}

// GetIPLocation returns the location for the specified IP.
func (c *Client) GetIPLocation(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "loc")
}

/* ORG */

// GetIPOrg returns the organization for the specified IP.
func GetIPOrg(ip net.IP) (string, error) {
	return DefaultClient.GetIPOrg(ip)
}

// GetIPOrg returns the organization for the specified IP.
func (c *Client) GetIPOrg(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "org")
}

/* POSTAL */

// GetIPPostal returns the postal for the specified IP.
func GetIPPostal(ip net.IP) (string, error) {
	return DefaultClient.GetIPPostal(ip)
}

// GetIPPostal returns the postal for the specified IP.
func (c *Client) GetIPPostal(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "postal")
}

/* TIMEZONE */

// GetIPTimezone returns the timezone for the specified IP.
func GetIPTimezone(ip net.IP) (string, error) {
	return DefaultClient.GetIPTimezone(ip)
}

// GetIPTimezone returns the timezone for the specified IP.
func (c *Client) GetIPTimezone(ip net.IP) (string, error) {
	return c.getIPSpecificStr(ip, "timezone")
}

func (c *Client) getIPSpecificStr(ip net.IP, spec string) (string, error) {
	v := new(string)
	res, err := c.getIPSpecific(ip, spec, v)
	if err != nil {
		return "", err
	}
	return *(res.(*string)), nil
}

func (c *Client) getIPSpecific(
	ip net.IP,
	spec string,
	data interface{},
) (interface{}, error) {
	var cacheKey string
	var relURL string

	if ip == nil {
		// NOTE: we assume that if no IP is given, the user has the same IP as
		// when a previous cache lookup happened, so if that result still
		// exists, we return it. This is an issue if the user's IP changes.
		cacheKey = "ip:" + spec + ":nil"
		relURL = spec
	} else {
		ipStr := ip.String()
		cacheKey = "ip:" + spec + ":" + ipStr
		relURL = ipStr + "/" + spec
	}

	// perform cache lookup.
	if c.Cache != nil {
		if res, err := c.Cache.Get(cacheKey); err == nil {
			return res, nil
		}
	}

	// prepare req
	req, err := c.NewRequest(relURL)
	if err != nil {
		return nil, err
	}

	// do req
	if _, err := c.Do(req, data); err != nil {
		return nil, err
	}

	// cache req result
	if c.Cache != nil {
		if err := c.Cache.Set(cacheKey, data); err != nil {
			return nil, err
		}
	}

	return data, nil
}
