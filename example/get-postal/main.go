package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	postal, err := ipinfo.GetIPPostal(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", postal)
}
