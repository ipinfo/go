# [<img src="https://ipinfo.io/static/ipinfo-small.svg" alt="IPinfo" width="24"/>](https://ipinfo.io/) IPinfo Go Client Library

[![License](http://img.shields.io/:license-apache-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/ipinfo/go-ipinfo/ipinfo?status.svg)](https://godoc.org/github.com/ipinfo/go-ipinfo/ipinfo)

This is the official Go client library for the [IPinfo.io](https://ipinfo.io) IP address API, allowing you to lookup your own IP address, or get any of the following details for an IP:
 - IP geolocation (city, region, country, postal code, latitude and longitude)
 - ASN details (ISP or network operator, associated domain name, and type, such as business, hosting or company)
 - Company details (the name and domain of the business that uses the IP address)
 - Carrier details (the name of the mobile carrier and MNC and MCC for that carrier if the IP is used exclusively for mobile traffic)

### Getting Started

You'll need an IPinfo API access token, which you can get by singing up for a free account at [https://ipinfo.io/signup](https://ipinfo.io/signup?ref=lib-go). 

The free plan is limited to 1,000 requests a day, and doesn't include some of the data fields such as IP type and company data. To enable all the data fields and additional request volumes see [https://ipinfo.io/pricing](https://ipinfo.io/pricing?ref=lib-go)

#### Installation

```
go get github.com/ipinfo/go-ipinfo/ipinfo
```

#### Quick Start

```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

func main() {
	// Get access token by signing up a free account at https://ipinfo.io/signup
	authTransport := ipinfo.AuthTransport{Token: "MY_TOKEN"}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)
	info, err := client.GetInfo(net.ParseIP("8.8.8.8"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
	// &{{8.8.8.8 Mountain View California US 37.3860,-122.0840 650 94035} google-public-dns-a.google.com }
}
```

### Other Libraries

There are official IPinfo client libraries available for many languages including PHP, Python, Go, Java, Ruby, and many popular frameworks such as Django, Rails and Laravel. There are also many third party libraries and integrations available for our API. 

https://ipinfo.io/developers/libraries

### About IPinfo

Founded in 2013, IPinfo prides itself on being the most reliable, accurate, and in-depth source of IP address data available anywhere. We process terabytes of data to produce our custom IP geolocation, company, carrier and IP type data sets. Our API handles over 12 billion requests a month for 100,000 businesses and developers.

[![image](https://avatars3.githubusercontent.com/u/15721521?s=128&u=7bb7dde5c4991335fb234e68a30971944abc6bf3&v=4)](https://ipinfo.io/)
