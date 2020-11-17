package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ipinfo/go/ipinfo"
	"github.com/ipinfo/go/ipinfo/cache"
)

func main() {
	ipinfo.SetCache(
		ipinfo.NewCache(
			cache.NewInMemory().WithExpiration(5 * time.Minute),
		),
	)
	ip := net.ParseIP("8.8.8.8")

	for i := 0; i < 2; i++ {
		fmt.Println([]string{"Actual requests", "From cache"}[i])
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
