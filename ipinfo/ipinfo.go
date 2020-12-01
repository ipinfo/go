//go:generate go run gen-fields.go

package ipinfo

import (
	"bytes"
	"net"
	"strings"
)

type Data struct {
	Ip       net.IP       `json:"ip"`
	Hostname string       `json:"hostname"`
	City     string       `json:"city"`
	Region   string       `json:"region"`
	Country  string       `json:"country"`
	Location string       `json:"loc"`
	Org      string       `json:"org"`
	Postal   string       `json:"postal"`
	Timezone string       `json:"timezone"`
	Asn      *DataAsn     `json:"asn"`
	Company  *DataCompany `json:"company"`
	Carrier  *DataCarrier `json:"carrier"`
	Privacy  *DataPrivacy `json:"privacy"`
	Abuse    *DataAbuse   `json:"abuse"`
	Domains  *DataDomains `json:"domains"`
}

type DataAsn struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

type DataCompany struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type DataCarrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

type DataPrivacy struct {
	Vpn     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Hosting bool `json:"hosting"`
}

type DataAbuse struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

type DataDomains struct {
	Ip      string   `json:"ip"`
	Total   uint64   `json:"total"`
	Domains []string `json:"domains"`
}

// Global, default client available to the user via `ipinfo.DefaultClient`.
var DefaultClient *Client

func init() {
	/* Create a global, default client. */
	DefaultClient = NewClient(nil, nil, "")
}
