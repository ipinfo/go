package ipinfo

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
)

// IPSummary is the full JSON response from the IP summary API.
type IPSummary struct {
	Total     uint64            `json:"total"`
	Unique    uint64            `json:"unique"`
	Countries map[string]uint64 `json:"countries"`
	Cities    map[string]uint64 `json:"cities"`
	Regions   map[string]uint64 `json:"regions"`
	ASNs      map[string]uint64 `json:"asns"`
	Companies map[string]uint64 `json:"companies"`
	IPTypes   map[string]uint64 `json:"ipTypes"`
	Routes    map[string]uint64 `json:"routes"`
	Carriers  map[string]uint64 `json:"carriers"`
	Mobile    uint64            `json:"mobile"`
	Domains   map[string]uint64 `json:"domains"`
	Privacy   struct {
		VPN     uint64 `json:"vpn"`
		Proxy   uint64 `json:"proxy"`
		Hosting uint64 `json:"hosting"`
		Tor     uint64 `json:"tor"`
	} `json:"privacy"`
	Anycast uint64 `json:"anycast"`
	Bogon   uint64 `json:"bogon"`
}

// GetIPSummary returns summarized results for a group of IPs.
//
// `ips` must contain at least 10 unique IPs, and a total maximum of 1000.
func GetIPSummary(ips []net.IP) (*IPSummary, error) {
	return DefaultClient.GetIPSummary(ips)
}

// GetIPSummary returns summarized results for a group of IPs.
//
// `ips` must contain at least 10 unique IPs, and a total maximum of 1000.
func (c *Client) GetIPSummary(ips []net.IP) (*IPSummary, error) {
	if len(ips) < 10 || len(ips) > 1000 {
		return nil, errors.New("unique ip count must be >10 && <1000")
	}

	jsonArrStr, err := json.Marshal(ips)
	if err != nil {
		return nil, err
	}
	jsonBuf := bytes.NewBuffer(jsonArrStr)

	req, err := c.newRequest(nil, "POST", "summarize", jsonBuf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	result := new(IPSummary)
	if _, err := c.do(req, result); err != nil {
		return nil, err
	}

	return result, nil
}
