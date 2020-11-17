package ipinfo // import "github.com/ipinfo/go/ipinfo"

import (
	"strings"
)

// ASNInfo represents ASN details.
type ASNInfo struct {
	ASN         string      `json:"asn"`
	Name        string      `json:"name"`
	Country     string      `json:"country"`
	Allocated   string      `json:"allocated"`
	Registry    string      `json:"registry"`
	Domain      string      `json:"domain"`
	NumberOfIPs uint64      `json:"num_ips"`
	Type        string      `json:"type"`
	Prefixes    []ASNPrefix `json:"prefixes"`
	Prefixes6   []ASNPrefix `json:"prefixes6"`
}

// ASNPrefix represents an ASN prefix.
type ASNPrefix struct {
	Netblock string `json:"netblock"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Country  string `json:"country"`
}

// InvalidASNError is reported when the invalid ASN was specified.
type InvalidASNError struct {
	ASN string
}

func (err *InvalidASNError) Error() string {
	return "invalid ASN: " + err.ASN
}

// ASN returns the details for the specified ASN.
func ASN(asn string) (*ASNInfo, error) {
	return c.ASN(asn)
}

// ASN returns the details for the specified ASN.
func (c *Client) ASN(asn string) (*ASNInfo, error) {
	if !strings.HasPrefix(asn, "AS") {
		return nil, &InvalidASNError{ASN: asn}
	}
	req, err := c.NewRequest(asn + "/json")
	if err != nil {
		return nil, err
	}
	v := new(ASNInfo)
	_, err = c.Do(req, v)
	return v, err
}
