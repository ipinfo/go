package ipinfo

import (
	"net"
)

type Core struct {
	Ip       net.IP       `json:"ip"`
	Hostname string       `json:"hostname"`
	City     string       `json:"city"`
	Region   string       `json:"region"`
	Country  string       `json:"country"`
	Location string       `json:"loc"`
	Org      string       `json:"org"`
	Postal   string       `json:"postal"`
	Timezone string       `json:"timezone"`
	Asn      *CoreAsn     `json:"asn"`
	Company  *CoreCompany `json:"company"`
	Carrier  *CoreCarrier `json:"carrier"`
	Privacy  *CorePrivacy `json:"privacy"`
	Abuse    *CoreAbuse   `json:"abuse"`
	Domains  *CoreDomains `json:"domains"`
}

type CoreAsn struct {
	Asn    string `json:"asn"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Route  string `json:"route"`
	Type   string `json:"type"`
}

type CoreCompany struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type CoreCarrier struct {
	Name string `json:"name"`
	Mcc  string `json:"mcc"`
	Mnc  string `json:"mnc"`
}

type CorePrivacy struct {
	Vpn     bool `json:"vpn"`
	Proxy   bool `json:"proxy"`
	Tor     bool `json:"tor"`
	Hosting bool `json:"hosting"`
}

type CoreAbuse struct {
	Address string `json:"address"`
	Country string `json:"country"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Network string `json:"network"`
	Phone   string `json:"phone"`
}

type CoreDomains struct {
	Ip      string   `json:"ip"`
	Total   uint64   `json:"total"`
	Domains []string `json:"domains"`
}
