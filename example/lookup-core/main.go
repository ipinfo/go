package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	client := ipinfo.NewCoreClient(nil, nil, os.Getenv("IPINFO_TOKEN"))

	ip := net.ParseIP("8.8.8.8")
	info, err := client.GetIPInfo(ip)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IP: %s\n", info.IP)
	if info.Geo != nil {
		fmt.Printf("City: %s\n", info.Geo.City)
		fmt.Printf("Region: %s\n", info.Geo.Region)
		fmt.Printf("Country: %s (%s)\n", info.Geo.Country, info.Geo.CountryCode)
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
	fmt.Printf("Anonymous: %v\n", info.IsAnonymous)
	fmt.Printf("Anycast: %v\n", info.IsAnycast)
	fmt.Printf("Hosting: %v\n", info.IsHosting)
	fmt.Printf("Mobile: %v\n", info.IsMobile)
	fmt.Printf("Satellite: %v\n", info.IsSatellite)
}
