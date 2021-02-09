package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	client := ipinfo.NewClient(nil, nil, os.Getenv("TOKEN"))
	result, err := client.GetIPSummary(
		[]net.IP{
			net.ParseIP("3.3.3.0"),
			net.ParseIP("3.3.3.1"),
			net.ParseIP("3.3.3.2"),
			net.ParseIP("3.3.3.3"),
			net.ParseIP("4.4.4.0"),
			net.ParseIP("4.4.4.1"),
			net.ParseIP("4.4.4.2"),
			net.ParseIP("4.4.4.3"),
			net.ParseIP("8.8.8.0"),
			net.ParseIP("8.8.8.1"),
			net.ParseIP("8.8.8.2"),
			net.ParseIP("8.8.8.3"),
			net.ParseIP("1.1.1.0"),
			net.ParseIP("1.1.1.1"),
			net.ParseIP("1.1.1.2"),
			net.ParseIP("1.1.1.3"),
			net.ParseIP("2.2.2.0"),
			net.ParseIP("2.2.2.1"),
			net.ParseIP("2.2.2.2"),
			net.ParseIP("2.2.2.3"),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result=%v\n", result)
}
