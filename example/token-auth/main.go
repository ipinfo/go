package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfoio/go-ipinfo/ipinfo"
)

func main() {
	var c *ipinfo.Client

	if token := os.Getenv("TOKEN"); token != "" {
		tp := ipinfo.AuthTransport{
			Token: token,
		}
		c = ipinfo.NewClient(tp.Client())
	} else {
		c = ipinfo.GetClient()
	}
	if len(os.Args) == 1 {
		info, err := c.GetInfo(nil)
		if err != nil {
			log.Println(err)
		}
		printInfo(info)
	}
	for _, s := range os.Args[1:] {
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}
		info, err := c.GetInfo(ip)
		if err != nil {
			log.Println(err)
		}
		printInfo(info)
	}
}

func printInfo(info *ipinfo.Info) {
	fmt.Printf("IP: %v\nHostname: %s\nOrganization: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nPhone: %s\nPostal: %s\n",
		info.IP, info.Hostname, info.Organization, info.City, info.Region, info.Country, info.Location, info.Phone, info.Postal)
}
