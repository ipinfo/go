# [<img src="https://ipinfo.io/static/ipinfo-small.svg" alt="IPinfo" width="24"/>](https://ipinfo.io/) IPinfo® Go Client Library

[![License](http://img.shields.io/:license-apache-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipinfo/go/v2/ipinfo.svg)](https://pkg.go.dev/github.com/ipinfo/go/v2/ipinfo)

This is the official Go client library for the [IPinfo.io](https://ipinfo.io) IP address API, allowing you to lookup your own IP address, or get any of the following details for other IP addresses:

- [IP to Geolocation](https://ipinfo.io/ip-geolocation-api) (city, region, country, postal code, latitude and longitude)
- [IP to ASN](https://ipinfo.io/asn-api) (ISP or network operator, associated domain name, and type, such as business, hosting or company)
- [IP to Company](https://ipinfo.io/ip-company-api) (the name and domain of the business that uses the IP address)
- [IP to Carrier](https://ipinfo.io/ip-carrier-api) (the name of the mobile carrier and MNC and MCC for that carrier if the IP is used exclusively for mobile traffic)

Check all the data we have for your IP address [here](https://ipinfo.io/what-is-my-ip).

### Getting Started

You'll need an IPinfo® API access token, which you can get by singing up for a free account at [https://ipinfo.io/signup](https://ipinfo.io/signup).

The free plan is limited to 50,000 requests per month, and doesn't include some of the data fields such as IP type and company data. To enable all the data fields and additional request volumes see [https://ipinfo.io/pricing](https://ipinfo.io/pricing)

#### Installation

```bash
go get github.com/ipinfo/go/v2/ipinfo
```

#### Quick Start

```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	// Get access token by signing up a free account at https://ipinfo.io/signup
	client := ipinfo.NewClient(nil, nil, "MY_TOKEN")
	info, err := client.GetIPInfo(net.ParseIP("8.8.8.8"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}
```

### Other Libraries

There are official [IPinfo® client libraries](https://ipinfo.io/developers/libraries) available for many languages including PHP, Python, Go, Java, Ruby, and many popular frameworks such as Django, Rails and Laravel. There are also many third party libraries and integrations available for our API.


### About IPinfo®

Founded in 2013, IPinfo prides itself on being the most reliable, accurate, and in-depth source of IP address data available anywhere. We process terabytes of data to produce our custom IP geolocation, company, carrier, VPN detection, hosted domains, and IP type data sets. Our API handles over 40 billion requests a month for 100,000 businesses and developers.

[![image](https://avatars3.githubusercontent.com/u/15721521?s=128&u=7bb7dde5c4991335fb234e68a30971944abc6bf3&v=4)](https://ipinfo.io/)
