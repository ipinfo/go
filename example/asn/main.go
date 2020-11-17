package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ipinfo/go/ipinfo"
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
	asnInfo, err := client.ASN("AS7922")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(asnInfo)
}
