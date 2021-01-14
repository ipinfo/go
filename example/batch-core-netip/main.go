package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
)

func main() {
	client := ipinfo.NewClient(
		nil,
		ipinfo.NewCache(cache.NewInMemory().WithExpiration(5*time.Minute)),
		os.Getenv("TOKEN"),
	)
	for i := 0; i < 3; i++ {
		fmt.Printf("doing lookup #%v\n", i)
		batchResult, err := client.GetIPInfoBatch(
			[]net.IP{
				net.ParseIP("1.1.1.1"),
				net.ParseIP("8.8.8.8"),
			},
			ipinfo.BatchReqOpts{
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
		fmt.Println()
	}
}
