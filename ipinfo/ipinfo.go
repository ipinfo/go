//go:generate go run gen-fields.go

package ipinfo

import (
	"bytes"
	"net"
	"strings"
)

var c *Client

func init() {
	c = NewClient(nil)
}

// Info represents full IP details.
type Info struct {
	Geo
	Hostname     string `json:"hostname"`
	Organization string `json:"org"`
}

// Geo represents IP geolocation information.
type Geo struct {
	IP       net.IP `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Phone    string `json:"phone"`
	Postal   string `json:"postal"`
}

// GetInfo returns full details for the specified IP. If nil was provieded
// instead of ip, it returns details for the caller's own IP.
func GetInfo(ip net.IP) (*Info, error) {
	return c.GetInfo(ip)
}

// GetInfo returns full details for the specified IP. If nil was provieded
// instead of ip, it returns details for the caller's own IP.
func (c *Client) GetInfo(ip net.IP) (*Info, error) {
	var s string
	if ip != nil {
		s = ip.String()
	}
	if c.Cache == nil {
		return c.requestInfo(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestInfo(s)
	})
	if err != nil {
		return nil, err
	}
	return v.(*Info), err
}

func (c *Client) requestInfo(s string) (*Info, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return nil, err
	}
	v := new(Info)
	_, err = c.Do(req, v)
	return v, err
}

// GetGeo returns geolocation information for the specified IP. If nil was provieded
// instead of ip, it returns details for the caller's own IP.
func GetGeo(ip net.IP) (*Geo, error) {
	return c.GetGeo(ip)
}

// GetGeo returns geolocation information for the specified IP. If nil was provieded
// instead of ip, it returns details for the caller's own IP.
func (c *Client) GetGeo(ip net.IP) (*Geo, error) {
	s := "geo"
	if ip != nil {
		s = ip.String() + "/geo"
	}
	if c.Cache == nil {
		return c.requestGeo(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestGeo(s)
	})
	if err != nil {
		return nil, err
	}
	return v.(*Geo), err
}

func (c *Client) requestGeo(s string) (*Geo, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return nil, err
	}
	v := new(Geo)
	_, err = c.Do(req, v)
	return v, err
}

// GetIP returns a specific field "ip" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func GetIP(ip net.IP) (net.IP, error) {
	return c.GetIP(ip)
}

// GetIP returns a specific field "ip" value from the
// API for the provided ip. If nil was provided instead of ip, it returns
// details for the caller's own IP.
func (c *Client) GetIP(ip net.IP) (net.IP, error) {
	s := "ip"
	if ip != nil {
		s = ip.String() + "/" + s
	}
	if c.Cache == nil {
		return c.requestIP(s)
	}
	v, err := c.Cache.GetOrRequest(s, func() (interface{}, error) {
		return c.requestIP(s)
	})
	if err != nil {
		return nil, err
	}
	return v.(net.IP), err
}

func (c *Client) requestIP(s string) (net.IP, error) {
	req, err := c.NewRequest(s)
	if err != nil {
		return nil, err
	}
	v := new(bytes.Buffer)
	_, err = c.Do(req, v)
	if err != nil {
		return nil, err
	}
	s = strings.TrimSpace(v.String())
	return net.ParseIP(s), nil
}
