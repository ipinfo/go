package ipinfo

import (
	"net"
)

// Core represents data from the Core API.
type Core struct {
	IP          net.IP       `json:"ip"`
	Hostname    string       `json:"hostname"`
	Anycast     bool         `json:"anycast"`
	City        string       `json:"city"`
	Region      string       `json:"region"`
	Country     string       `json:"country"`
	CountryName string       `json:'-"`
	Location    string       `json:"loc"`
	Org         string       `json:"org"`
	Postal      string       `json:"postal"`
	Timezone    string       `json:"timezone"`
	ASN         *CoreASN     `json:"asn"`
	Company     *CoreCompany `json:"company"`
	Carrier     *CoreCarrier `json:"carrier"`
	Privacy     *CorePrivacy `json:"privacy"`
	Abuse       *CoreAbuse   `json:"abuse"`
	Domains     *CoreDomains `json:"domains"`
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
	MCC  string `json:"mcc"`
	MNC  string `json:"mnc"`
}

// CorePrivacy represents privacy data for the Core API.
type CorePrivacy struct {
	VPN     bool `json:"vpn"`
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

// Set `v.CountryName` properly by mapping country abbreviation to full country
// name.
func (v *Core) setCountryName() {
	if v.Country != "" {
		v.CountryName = countriesMap[v.Country]
	}
}

/* CORE */

// GetIPInfo returns the details for the specified IP.
func GetIPInfo(ip net.IP) (*Core, error) {
	return DefaultClient.GetIPInfo(ip)
}

// GetIPInfo returns the details for the specified IP.
func (c *Client) GetIPInfo(ip net.IP) (*Core, error) {
	relURL := ""
	if ip != nil {
		relURL = ip.String()
	}

	// perform cache lookup.
	if c.Cache != nil {
		if res, err := c.Cache.Get(relURL); err == nil {
			return res.(*Core), nil
		}
	}

	// prepare req
	req, err := c.newRequest(nil, "GET", relURL, nil)
	if err != nil {
		return nil, err
	}

	// do req
	v := new(Core)
	if _, err := c.do(req, v); err != nil {
		return nil, err
	}

	// format
	v.setCountryName()

	// cache req result
	if c.Cache != nil {
		if err := c.Cache.Set(relURL, v); err != nil {
			// NOTE: still return the value even if the cache fails.
			return v, err
		}
	}

	return v, nil
}

/* IP ADDRESS */

// GetIPAddr returns the IP address that IPinfo sees when you make a request.
func GetIPAddr() (string, error) {
	return DefaultClient.GetIPAddr()
}

// GetIPAddr returns the IP address that IPinfo sees when you make a request.
func (c *Client) GetIPAddr() (string, error) {
	core, err := c.GetIPInfo(nil)
	if err != nil {
		return "", err
	}
	return core.IP.String(), nil
}

/* HOSTNAME */

// GetIPHostname returns the hostname of the domain on the specified IP.
func GetIPHostname(ip net.IP) (string, error) {
	return DefaultClient.GetIPHostname(ip)
}

// GetIPHostname returns the hostname of the domain on the specified IP.
func (c *Client) GetIPHostname(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Hostname, nil
}

/* CITY */

// GetIPCity returns the city for the specified IP.
func GetIPCity(ip net.IP) (string, error) {
	return DefaultClient.GetIPCity(ip)
}

// GetIPCity returns the city for the specified IP.
func (c *Client) GetIPCity(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.City, nil
}

/* REGION */

// GetIPRegion returns the region for the specified IP.
func GetIPRegion(ip net.IP) (string, error) {
	return DefaultClient.GetIPRegion(ip)
}

// GetIPRegion returns the region for the specified IP.
func (c *Client) GetIPRegion(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Region, nil
}

/* COUNTRY */

// GetIPCountry returns the country for the specified IP.
func GetIPCountry(ip net.IP) (string, error) {
	return DefaultClient.GetIPCountry(ip)
}

// GetIPCountry returns the country for the specified IP.
func (c *Client) GetIPCountry(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Country, nil
}

/* COUNTRY NAME */

// GetIPCountryName returns the full country name for the specified IP.
func GetIPCountryName(ip net.IP) (string, error) {
	return DefaultClient.GetIPCountryName(ip)
}

// GetIPCountryName returns the full country name for the specified IP.
func (c *Client) GetIPCountryName(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.CountryName, nil
}

/* LOCATION */

// GetIPLocation returns the location for the specified IP.
func GetIPLocation(ip net.IP) (string, error) {
	return DefaultClient.GetIPLocation(ip)
}

// GetIPLocation returns the location for the specified IP.
func (c *Client) GetIPLocation(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Location, nil
}

/* ORG */

// GetIPOrg returns the organization for the specified IP.
func GetIPOrg(ip net.IP) (string, error) {
	return DefaultClient.GetIPOrg(ip)
}

// GetIPOrg returns the organization for the specified IP.
func (c *Client) GetIPOrg(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Org, nil
}

/* POSTAL */

// GetIPPostal returns the postal for the specified IP.
func GetIPPostal(ip net.IP) (string, error) {
	return DefaultClient.GetIPPostal(ip)
}

// GetIPPostal returns the postal for the specified IP.
func (c *Client) GetIPPostal(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Postal, nil
}

/* TIMEZONE */

// GetIPTimezone returns the timezone for the specified IP.
func GetIPTimezone(ip net.IP) (string, error) {
	return DefaultClient.GetIPTimezone(ip)
}

// GetIPTimezone returns the timezone for the specified IP.
func (c *Client) GetIPTimezone(ip net.IP) (string, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return "", err
	}
	return core.Timezone, nil
}

/* ASN */

// GetIPASN returns the ASN details for the specified IP.
func GetIPASN(ip net.IP) (*CoreASN, error) {
	return DefaultClient.GetIPASN(ip)
}

// GetIPASN returns the ASN details for the specified IP.
func (c *Client) GetIPASN(ip net.IP) (*CoreASN, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.ASN, nil
}

/* COMPANY */

// GetIPCompany returns the company details for the specified IP.
func GetIPCompany(ip net.IP) (*CoreCompany, error) {
	return DefaultClient.GetIPCompany(ip)
}

// GetIPCompany returns the company details for the specified IP.
func (c *Client) GetIPCompany(ip net.IP) (*CoreCompany, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.Company, nil
}

/* CARRIER */

// GetIPCarrier returns the carrier details for the specified IP.
func GetIPCarrier(ip net.IP) (*CoreCarrier, error) {
	return DefaultClient.GetIPCarrier(ip)
}

// GetIPCarrier returns the carrier details for the specified IP.
func (c *Client) GetIPCarrier(ip net.IP) (*CoreCarrier, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.Carrier, nil
}

/* PRIVACY */

// GetIPPrivacy returns the privacy details for the specified IP.
func GetIPPrivacy(ip net.IP) (*CorePrivacy, error) {
	return DefaultClient.GetIPPrivacy(ip)
}

// GetIPPrivacy returns the privacy details for the specified IP.
func (c *Client) GetIPPrivacy(ip net.IP) (*CorePrivacy, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.Privacy, nil
}

/* ABUSE */

// GetIPAbuse returns the abuse details for the specified IP.
func GetIPAbuse(ip net.IP) (*CoreAbuse, error) {
	return DefaultClient.GetIPAbuse(ip)
}

// GetIPAbuse returns the abuse details for the specified IP.
func (c *Client) GetIPAbuse(ip net.IP) (*CoreAbuse, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.Abuse, nil
}

/* DOMAINS */

// GetIPDomains returns the domains details for the specified IP.
func GetIPDomains(ip net.IP) (*CoreDomains, error) {
	return DefaultClient.GetIPDomains(ip)
}

// GetIPDomains returns the domains details for the specified IP.
func (c *Client) GetIPDomains(ip net.IP) (*CoreDomains, error) {
	core, err := c.GetIPInfo(ip)
	if err != nil {
		return nil, err
	}
	return core.Domains, nil
}
