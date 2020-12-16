package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	// Get access token by signing up a free account at
	// https://ipinfo.io/signup.
	// Provide token as an environment variable `TOKEN`,
	// e.g. TOKEN="XXXXXXXXXXXXXX" go run main.go
	client := ipinfo.NewClient(nil, nil, os.Getenv("TOKEN"))
	asnDetails, err := client.GetASNDetails("AS7922")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", asnDetails)
}
