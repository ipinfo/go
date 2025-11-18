package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	client := ipinfo.NewPlusClient(nil, nil, os.Getenv("IPINFO_TOKEN"))

	ip := net.ParseIP("8.8.8.8")
	info, err := client.GetIPInfo(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP: %s\n", info.IP)
	fmt.Printf("Hostname: %s\n", info.Hostname)
	if info.Geo != nil {
		fmt.Printf("City: %s\n", info.Geo.City)
		fmt.Printf("Region: %s (%s)\n", info.Geo.Region, info.Geo.RegionCode)
		fmt.Printf("Country: %s (%s)\n", info.Geo.Country, info.Geo.CountryCode)
		fmt.Printf("Continent: %s (%s)\n", info.Geo.Continent, info.Geo.ContinentCode)
		fmt.Printf("Location: %f, %f\n", info.Geo.Latitude, info.Geo.Longitude)
		fmt.Printf("Timezone: %s\n", info.Geo.Timezone)
		fmt.Printf("Postal Code: %s\n", info.Geo.PostalCode)
	}
	if info.AS != nil {
		fmt.Printf("ASN: %s\n", info.AS.ASN)
		fmt.Printf("AS Name: %s\n", info.AS.Name)
		fmt.Printf("AS Domain: %s\n", info.AS.Domain)
		fmt.Printf("AS Type: %s\n", info.AS.Type)
	}
	if info.Mobile != nil {
		fmt.Printf("Mobile - Name: %s, MCC: %s, MNC: %s\n",
			info.Mobile.Name, info.Mobile.MCC, info.Mobile.MNC)
	}
	if info.Anonymous != nil {
		fmt.Printf("Anonymous - Proxy: %v, Relay: %v, Tor: %v, VPN: %v\n",
			info.Anonymous.IsProxy, info.Anonymous.IsRelay, info.Anonymous.IsTor, info.Anonymous.IsVPN)
	}
	fmt.Printf("Anonymous: %v\n", info.IsAnonymous)
	fmt.Printf("Anycast: %v\n", info.IsAnycast)
	fmt.Printf("Hosting: %v\n", info.IsHosting)
	fmt.Printf("Mobile: %v\n", info.IsMobile)
	fmt.Printf("Satellite: %v\n", info.IsSatellite)
	if info.Abuse != nil {
		fmt.Printf("Abuse Contact - Email: %s, Name: %s\n", info.Abuse.Email, info.Abuse.Name)
	}
	if info.Company != nil {
		fmt.Printf("Company - Name: %s, Domain: %s, Type: %s\n",
			info.Company.Name, info.Company.Domain, info.Company.Type)
	}
	if info.Privacy != nil {
		fmt.Printf("Privacy - VPN: %v, Proxy: %v, Tor: %v, Relay: %v, Hosting: %v\n",
			info.Privacy.VPN, info.Privacy.Proxy, info.Privacy.Tor, info.Privacy.Relay, info.Privacy.Hosting)
	}
	if info.Domains != nil {
		fmt.Printf("Domains - Total: %d\n", info.Domains.Total)
		if len(info.Domains.Domains) > 0 {
			maxDomains := 3
			if len(info.Domains.Domains) < maxDomains {
				maxDomains = len(info.Domains.Domains)
			}
			fmt.Printf("Sample Domains: %v\n", info.Domains.Domains[:maxDomains])
		}
	}
}
