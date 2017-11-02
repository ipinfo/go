package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfoio/go-ipinfo/ipinfo"
)

func main() {
	info, err := ipinfo.GetInfo(net.ParseIP("2a03:2880:f10a:83:face:b00c:0:25de"))
	if err != nil {
		log.Fatal(err)
	}
	printInfo(info)
}

func printInfo(info *ipinfo.Info) {
	fmt.Printf("IP: %v\nHostname: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPhone: %s\nPostal: %s\n",
		info.IP, info.Hostname, info.Organization, info.City, info.Region, info.Country, info.Location, info.Phone, info.Postal)
}
