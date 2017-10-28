package main

import (
	"fmt"
	"log"

	"github.com/ipinfoio/go-ipinfo/ipinfo"
)

func main() {
	info, err := ipinfo.GetInfo(nil)
	if err != nil {
		log.Fatal(err)
	}
	printInfo(info)
}

func printInfo(info *ipinfo.Info) {
	fmt.Printf("IP: %v\nHostname: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPhone: %s\nPostal: %s\n",
		info.IP, info.Hostname, info.Organization, info.City, info.Region, info.Country, info.Location, info.Phone, info.Postal)
}
