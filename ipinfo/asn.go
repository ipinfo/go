package ipinfo

import (
	"strings"
)

type ASNDetails struct {
	ASN         string             `json:"asn"`
	Name        string             `json:"name"`
	Country     string             `json:"country"`
	Allocated   string             `json:"allocated"`
	Registry    string             `json:"registry"`
	Domain      string             `json:"domain"`
	NumIPs      uint64             `json:"num_ips"`
	Type        string             `json:"type"`
	Prefixes    []ASNDetailsPrefix `json:"prefixes"`
	Prefixes6   []ASNDetailsPrefix `json:"prefixes6"`
	Peers       []string           `json:"peers"`
	Upstreams   []string           `json:"upstreams"`
	Downstreams []string           `json:"downstreams"`
}

type ASNDetailsPrefix struct {
	Netblock string `json:"netblock"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Size     string `json:"size"`
	Status   string `json:"status"`
	Domain   string `json:"domain"`
}

// InvalidASNError is reported when the invalid ASN was specified.
type InvalidASNError struct {
	ASN string
}

func (err *InvalidASNError) Error() string {
	return "invalid ASN: " + err.ASN
}

// GetASNDetails returns the details for the specified ASN.
func GetASNDetails(asn string) (*ASNDetails, error) {
	return DefaultClient.GetASNDetails(asn)
}

// GetASNDetails returns the details for the specified ASN.
func (c *Client) GetASNDetails(asn string) (*ASNDetails, error) {
	if !strings.HasPrefix(asn, "AS") {
		return nil, &InvalidASNError{ASN: asn}
	}
	req, err := c.NewRequest(asn + "/json")
	if err != nil {
		return nil, err
	}
	v := new(ASNDetails)
	_, err = c.Do(req, v)
	return v, err
}
