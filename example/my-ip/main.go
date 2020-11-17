package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	info, err := ipinfo.GetIP(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
}
