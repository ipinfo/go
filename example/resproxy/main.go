package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	// Get access token by signing up a free account at
	// https://ipinfo.io/signup.
	// Provide token as an environment variable `TOKEN`,
	// e.g. TOKEN="XXXXXXXXXXXXXX" go run main.go
	client := ipinfo.NewClient(nil, nil, os.Getenv("TOKEN"))
	resproxy, err := client.GetResproxy("175.107.211.204")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", resproxy)
}
