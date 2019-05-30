package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

func main() {
	// Get access token by signing up a free account at https://ipinfo.io/signup
	// Provide token as an environment variable `TOKEN`,
	// e.g. TOKEN="XXXXXXXXXXXXXX" go run main.go
	authTransport := ipinfo.AuthTransport{
		Token: os.Getenv("TOKEN"),
	}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)
	info, err := client.GetInfo(net.ParseIP("8.8.8.8"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
	// &{{8.8.8.8 Mountain View California US 37.3860,-122.0840 650 94035} google-public-dns-a.google.com }
}
