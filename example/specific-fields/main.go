package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	if len(os.Args) == 1 {
		if v, err := ipinfo.GetIP(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("IP:", v)
		}

		if v, err := ipinfo.GetHostname(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Hostname:", v)
		}

		if v, err := ipinfo.GetOrganization(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Organization:", v)
		}

		if v, err := ipinfo.GetCity(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("City:", v)
		}

		if v, err := ipinfo.GetRegion(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Region:", v)
		}

		if v, err := ipinfo.GetCountry(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Country:", v)
		}

		if v, err := ipinfo.GetLocation(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Location:", v)
		}

		if v, err := ipinfo.GetPhone(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Phone:", v)
		}

		if v, err := ipinfo.GetPostal(nil); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Postal:", v)
		}
	}
	for _, s := range os.Args[1:] {
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}
		if v, err := ipinfo.GetIP(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("IP:", v)
		}

		if v, err := ipinfo.GetHostname(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Hostname:", v)
		}

		if v, err := ipinfo.GetOrganization(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Organization:", v)
		}

		if v, err := ipinfo.GetCity(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("City:", v)
		}

		if v, err := ipinfo.GetRegion(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Region:", v)
		}

		if v, err := ipinfo.GetCountry(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Country:", v)
		}

		if v, err := ipinfo.GetLocation(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Location:", v)
		}

		if v, err := ipinfo.GetPhone(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Phone:", v)
		}

		if v, err := ipinfo.GetPostal(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Postal:", v)
		}
	}
}
