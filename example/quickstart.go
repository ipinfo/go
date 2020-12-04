package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	// Get access token by signing up a free account at
	// https://ipinfo.io/signup.
	// Provide token as an environment variable `TOKEN`,
	// e.g. TOKEN="XXXXXXXXXXXXXX" go run main.go
	client := ipinfo.NewClient(nil, nil, os.Getenv("TOKEN"))
	core, err := client.GetIPInfo(net.ParseIP("8.8.8.8"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", core)
}
