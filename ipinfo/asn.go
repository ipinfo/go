package ipinfo

import (
	"strings"
)

type AsnDetails struct {
	Asn         string             `json:"asn"`
	Name        string             `json:"name"`
	Country     string             `json:"country"`
	Allocated   string             `json:"allocated"`
	Registry    string             `json:"registry"`
	Domain      string             `json:"domain"`
	NumIps      uint64             `json:"num_ips"`
	Type        string             `json:"type"`
	Prefixes    []AsnDetailsPrefix `json:"prefixes"`
	Prefixes6   []AsnDetailsPrefix `json:"prefixes6"`
	Peers       []string           `json:"peers"`
	Upstreams   []string           `json:"upstreams"`
	Downstreams []string           `json:"downstreams"`
}

type AsnDetailsPrefix struct {
	Netblock string `json:"netblock"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Size     string `json:"size"`
	Status   string `json:"status"`
	Domain   string `json:"domain"`
}

// InvalidAsnError is reported when the invalid ASN was specified.
type InvalidAsnError struct {
	Asn string
}

func (err *InvalidAsnError) Error() string {
	return "invalid ASN: " + err.ASN
}

// AsnDetails returns the details for the specified ASN.
func AsnDetails(asn string) (*AsnDetails, error) {
	return DefaultClient.AsnDetails(asn)
}

// AsnDetails returns the details for the specified ASN.
func (c *Client) AsnDetails(asn string) (*AsnDetails, error) {
	if !strings.HasPrefix(asn, "AS") {
		return nil, &InvalidAsnError{Asn: asn}
	}
	req, err := c.NewRequest(asn + "/json")
	if err != nil {
		return nil, err
	}
	v := new(AsnDetails)
	_, err = c.Do(req, v)
	return v, err
}
