# [<img src="https://ipinfo.io/static/ipinfo-small.svg" alt="IPinfo" width="24"/>](https://ipinfo.io/) IPinfo Go Client Library

[![License](http://img.shields.io/:license-apache-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipinfo/go/v2/ipinfo.svg)](https://pkg.go.dev/github.com/ipinfo/go/v2/ipinfo)

This is the official Go client library for the [IPinfo.io](https://ipinfo.io) IP address API, allowing you to look up your own IP address, or get any of the following details for other IP addresses:

- [IP to Geolocation](https://ipinfo.io/ip-geolocation-api) (city, region, country, postal code, latitude, and longitude)
- [IP to ASN](https://ipinfo.io/asn-api) (ISP or network operator, associated domain name, and type, such as business, hosting, or company)
- [IP to Company](https://ipinfo.io/ip-company-api) (the name and domain of the business that uses the IP address)
- [IP to Carrier](https://ipinfo.io/ip-carrier-api) (the name of the mobile carrier and MNC and MCC for that carrier if the IP is used exclusively for mobile traffic)

Check all the data we have for your IP address [here](https://ipinfo.io/what-is-my-ip).


- [Getting Started](#getting-started)
	- [Installation](#installation)
	- [Quickstart](#quickstart)
- [Authentication](#authentication)
- [Internationalization](#internationalization)
	- [Country Name](#country-name)
	- [European Union (EU) Country](#european-union-eu-country)
	- [Country Flag](#country-flag)
	- [Country Currency](#country-currency)
	- [Continent](#continent)
- [Map IP Address](#map-ip-address)
- [Summarize IP Address](#summarize-ip-address)
- [Caching](#caching)
- [Batch Operations / Bulk Lookup](#batch-operations--bulk-lookup)
- [Other Libraries](#other-libraries)
- [About IPinfo](#about-ipinfo)

# Getting Started


You'll need an IPinfo API access token, which you can get by signing up for a free account at [https://ipinfo.io/signup](https://ipinfo.io/signup).

The free plan is limited to 50,000 requests per month, and doesn't include some of the data fields such as IP type and company data. To enable all the data fields and additional request volumes see [https://ipinfo.io/pricing](https://ipinfo.io/pricing)

You can find the full package-level documentation here: https://pkg.go.dev/github.com/ipinfo/go/v2/ipinfo

## Installation

```bash
go get github.com/ipinfo/go/v2/ipinfo
```

## Quickstart

Basic usage of the package.


```go
package main

import (
	"fmt"
	"log"
	"net"
	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	const token = "YOUR_TOKEN"
	
	// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
	client := ipinfo.NewClient(nil, nil, token)

	const ip_address = "8.8.8.8"
	info, err := client.GetIPInfo(net.ParseIP(ip_address))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
	// Output: {8.8.8.8 dns.google false true Mountain View California US United States...
}
```

This data is available even on our free tier which includes up to 50,000 IP geolocation requests per month.

# Authentication

The IPinfo Go library can be authenticated with your IPinfo API access token, which is passed as the third positional argument of the `ipinfo.NewClient()` method. Your IPInfo access token can be found in the account section of IPinfo's website after you have signed in: https://ipinfo.io/account/token

```go
const token = "YOUR_TOKEN"
// params: httpClient, cache, token. `http.DefaultClient` and no cache will be used in case of `nil`.
client := ipinfo.NewClient(nil, nil, token)
```

# Internationalization

## Country Name

`info.Country` returns the  ISO 3166 country code and `info.CountryName` returns the entire conuntry name:

```go
fmt.Println(info.Country)
// Output: US
fmt.Println(info.CountryName)
// Output: United States
```

## European Union (EU) Country

`info.IsEU` returns a boolean response to see if a country is a European Union country or not.

```go
fmt.Println(info.IsEU)
// Output: false
```

## Country Flag

Get country flag as an emoji and its Unicode value with `info.CountryFlag.Emoji` and `info.CountryFlag.Unicode` respectively.

```go
fmt.Println(info.CountryFlag.Emoji)
// Output: 🇳🇿 
fmt.Println(info.CountryFlag.Unicode)
// Output: "U+1F1F3 U+1F1FF"
```

## Country Flag URL

Get the link of a country's flag image.

```go
fmt.Println(info.CountryFlagURL)
// Output: https://cdn.ipinfo.io/static/images/countries-flags/US.svg"
```

## Country Currency

Get country's currency code and its symbol with `info.CountryCurrency.Code` and `info.CountryCurrency.Symbol` respectively.

```go
fmt.Println(info.CountryCurrency.Code)
// Output: USD 
fmt.Println(info.CountryCurrency.Symbol)
// Output: $
```

## Continent

Get the IP's continent code and its name with `info.Continent.Code` and `info.Continent.Name` respectively.

```go
fmt.Println(info.Continent.Code)
// Output: NA 
fmt.Println(info.Continent.Name)
// Output: North America
```

# Map IP Address

You can map up to 500,000 IP addresses all at once using the `GetIPMap` command. You can input:

- IP addresses (IPV4 and IPV6 both)
- IP Ranges or Netblock
- ASN

After the operation, you will be presented with a URL to a map generated on the IPinfo website.

IP Map Code:

```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	client := ipinfo.NewClient(nil, nil, "YOUR_TOKEN")
	result, err := client.GetIPMap(
		[]net.IP{
			net.ParseIP("136.111.157.61"),
			net.ParseIP("231.163.78.134"),
			// ...
			net.ParseIP("228.128.213.179"),
			net.ParseIP("103.172.175.76"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

```

Result:

See the output example map: https://ipinfo.io/tools/map/f27c7d40-3ff0-4ac2-878f-8d953dbcd3c8

# Summarize IP Address

Summarize IP addresses with `GetIPSummary` and output a report. 

```go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	client := ipinfo.NewClient(nil, nil, "YOUR_TOKEN")
	result, err := client.GetIPSummary(
		[]net.IP{
			net.ParseIP("171.164.236.38"),
			net.ParseIP("206.132.224.214"),
			// ....
			net.ParseIP("208.191.89.104"),
			net.ParseIP("81.216.14.76"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	// Ouptut: {100 100 map[CN:7 DE:4 JP:12 MX:3 US:32] map[Columbus, ...

}

```

# Batch Operations / Bulk Lookup

You can do batch lookups or bulk lookups quite easily as well. The inputs supported:

- IP addresses. IPV4 and IPV6 both
- ASN
- Specific field endpoint of an IP address e.g. `8.8.8.8/country`
   

```go
package main

import (
	"fmt"
	"log"
	"time"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
)

func main() {
	client := ipinfo.NewClient(
		nil,
		ipinfo.NewCache(cache.NewInMemory().WithExpiration(5*time.Minute)),
		"YOUR_TOKEN",
	)

	// batchResult will contain all the batch lookup data
	batchResult, err := client.GetBatch(
		[]string{
			"104.193.114.182",                    // you can pass IPV4 address
			"8.8.8.8/country",                    // you can get specific information
			"AS36811",                            // you can lookup ASN details
			"2a03:2880:f10a:83:face:b00c:0:25de", // IPV6 address
		},
		ipinfo.BatchReqOpts{
			BatchSize:       2,
			TimeoutPerBatch: 0,
			TimeoutTotal:    5,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range batchResult {
		fmt.Printf("k=%v v=%v\n", k, v)
	}
}
```

Examples of Batch / Bulk Lookup:

- [Batch ASN](/example/batch-asn)
- [Batch Core Net-IP](/example/batch-core-netip)
- [Batch Core str](/example/batch-core-str)
- [Batch Generic](/example/batch-generic)

The loop declaration in the batch lookup showcases the "caching" capability of the IPinfo package.

# Other Libraries

There are official [IPinfo client libraries](https://ipinfo.io/developers/libraries) available for many languages including PHP, Python, Go, Java, Ruby, and many popular frameworks such as Django, Rails, and Laravel. There are also many third-party libraries and integrations available for our API.

# About IPinfo

Founded in 2013, IPinfo prides itself on being the most reliable, accurate, and in-depth source of IP address data available anywhere. We process terabytes of data to produce our custom IP geolocation, company, carrier, VPN detection, hosted domains, and IP type data sets. Our API handles over 40 billion requests a month for 100,000 businesses and developers.

[![image](https://avatars3.githubusercontent.com/u/15721521?s=128&u=7bb7dde5c4991335fb234e68a30971944abc6bf3&v=4)](https://ipinfo.io/)
