package ipinfo

// ResproxyDetails represents residential proxy detection details for an IP.
type ResproxyDetails struct {
	IP              string  `json:"ip"`
	LastSeen        string  `json:"last_seen"`
	PercentDaysSeen float64 `json:"percent_days_seen"`
	Service         string  `json:"service"`
}

// GetResproxy returns the residential proxy details for the specified IP.
func GetResproxy(ip string) (*ResproxyDetails, error) {
	return DefaultClient.GetResproxy(ip)
}

// GetResproxy returns the residential proxy details for the specified IP.
func (c *Client) GetResproxy(ip string) (*ResproxyDetails, error) {
	// perform cache lookup.
	cacheKey := cacheKey("resproxy:" + ip)
	if c.Cache != nil {
		if res, err := c.Cache.Get(cacheKey); err == nil {
			return res.(*ResproxyDetails), nil
		}
	}

	// prepare req
	req, err := c.newRequest(nil, "GET", "resproxy/"+ip, nil)
	if err != nil {
		return nil, err
	}

	// do req
	v := new(ResproxyDetails)
	if _, err := c.do(req, v); err != nil {
		return nil, err
	}

	// cache req result
	if c.Cache != nil {
		if err := c.Cache.Set(cacheKey, v); err != nil {
			return v, err
		}
	}

	return v, nil
}
