package ipinfo

import (
	"bytes"
	"context"
	"encoding/json"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	batchMaxSize           = 1000
	batchReqTimeoutDefault = 5
)

// Internal batch type used by common batch functionality to temporarily store
// the URL-to-result mapping in a half-decoded state (specifically the value
// not being decoded yet). This allows us to decode the value to a proper
// concrete type like `Core` or `ASNDetails` after analyzing the key to
// determine which one it should be.
type batch map[string]json.RawMessage

// Batch is a mapped result of any valid API endpoint (e.g. `<ip>`,
// `<ip>/<field>`, `<asn>`, etc) to its corresponding data.
//
// The corresponding value will be either `*Core`, `*ASNDetails` or a generic
// map for unknown value results.
type Batch map[string]interface{}

// BatchCore is a mapped result of IPs to their corresponding `Core` data.
type BatchCore map[string]*Core

// BatchASNDetails is a mapped result of ASNs to their corresponding
// `ASNDetails` data.
type BatchASNDetails map[string]*ASNDetails

// BatchReqOpts are options input into batch request functions.
type BatchReqOpts struct {
	// BatchSize is the internal batch size used per API request; the IPinfo
	// API has a maximum batch size, but the batch request functions available
	// in this library do not. Therefore the library chunks the input slices
	// internally into chunks of size `BatchSize`, clipping to the maximum
	// allowed by the IPinfo API.
	//
	// 0 means to use the default batch size which is the max allowed by the
	// IPinfo API.
	BatchSize uint32

	// TimeoutPerBatch is the timeout in seconds that each batch of size
	// `BatchSize` will have for its own request.
	//
	// 0 means to use a default of 5 seconds; any negative number will turn it
	// off; turning it off does _not_ disable the effects of `TimeoutTotal`.
	TimeoutPerBatch int64

	// TimeoutTotal is the total timeout in seconds for all batch requests in a
	// batch request function to complete.
	//
	// 0 means no total timeout; `TimeoutPerBatch` will still apply.
	TimeoutTotal uint64

	// Filter, if turned on, will filter out a URL whose value was deemed empty
	// on the server.
	Filter bool
}

/* GENERIC */

// GetBatch does a batch request for all `urls` at once.
func GetBatch(
	urls []string,
	opts BatchReqOpts,
) (Batch, error) {
	return DefaultClient.GetBatch(urls, opts)
}

// GetBatch does a batch request for all `urls` at once.
func (c *Client) GetBatch(
	urls []string,
	opts BatchReqOpts,
) (Batch, error) {
	var batchSize int
	var lookupUrls []string
	var result Batch
	var wg sync.WaitGroup
	var mu sync.Mutex

	// use correct batch size; default/clip to `batchMaxSize`.
	if opts.BatchSize == 0 || opts.BatchSize > batchMaxSize {
		batchSize = batchMaxSize
	} else {
		batchSize = int(opts.BatchSize)
	}

	// if the cache is available, filter out URLs already cached.
	result = make(Batch, len(urls))
	if c.Cache != nil {
		lookupUrls = make([]string, 0, len(urls)/2)
		for _, url := range urls {
			if res, err := c.Cache.Get(url); err == nil {
				result[url] = res
			} else {
				lookupUrls = append(lookupUrls, url)
			}
		}
	} else {
		lookupUrls = urls
	}

	// everything cached.
	if len(lookupUrls) == 0 {
		return result, nil
	}

	for i := 0; i < len(lookupUrls); i += batchSize {
		end := i + batchSize
		if end > len(lookupUrls) {
			end = len(lookupUrls)
		}

		wg.Add(1)
		go func(urlsChunk []string) {
			var postURL string
			var timeoutPerBatch int64

			defer wg.Done()

			// TODO manage errors and timeouts properly.
			// TODO total timeout.

			if opts.TimeoutPerBatch == 0 {
				timeoutPerBatch = batchReqTimeoutDefault
			} else {
				timeoutPerBatch = opts.TimeoutPerBatch
			}

			// prepare request.

			ctx, cancel := context.WithTimeout(
				context.Background(),
				time.Duration(timeoutPerBatch)*time.Second,
			)
			defer cancel()

			jsonArrStr, err := json.Marshal(urlsChunk)
			if err != nil {
				return
			}

			if opts.Filter {
				postURL = "batch?filter=1"
			} else {
				postURL = "batch"
			}

			jsonBuf := bytes.NewBuffer(jsonArrStr)

			req, err := c.newRequest(ctx, "POST", postURL, jsonBuf)
			if err != nil {
				return
			}
			req.Header.Set("Content-Type", "application/json")

			// temporarily make a new local result map so that we can read the
			// network data into it; once we have it local we'll merge it with
			// `result` in a concurrency-safe way.
			localResult := new(batch)
			if _, err := c.do(req, localResult); err != nil {
				return
			}

			// update final result.
			mu.Lock()
			for k, v := range *localResult {
				if strings.HasPrefix(k, "AS") {
					decodedV := new(ASNDetails)
					if err := json.Unmarshal(v, decodedV); err != nil {
						return
					}

					decodedV.setCountryName()
					result[k] = decodedV
				} else if net.ParseIP(k) != nil {
					decodedV := new(Core)
					if err := json.Unmarshal(v, decodedV); err != nil {
						return
					}

					decodedV.setCountryName()
					result[k] = decodedV
				} else {
					decodedV := new(interface{})
					if err := json.Unmarshal(v, decodedV); err != nil {
						return
					}

					result[k] = decodedV
				}
			}
			mu.Unlock()
		}(lookupUrls[i:end])

	}
	wg.Wait()

	// we delay inserting into the cache until now because:
	// 1. it's likely more cache-line friendly.
	// 2. doing it while updating `result` inside the request workers would be
	//    problematic if the cache is external since we take a mutex lock for
	//    that entire period.
	if c.Cache != nil {
		for _, url := range lookupUrls {
			if v, exists := result[url]; exists {
				if err := c.Cache.Set(url, v); err != nil {
					// NOTE: still return the result even if the cache fails.
					return result, err
				}
			}
		}
	}

	return result, nil
}

/* CORE */

// GetIPInfoBatch does a batch request for all `ips` at once.
func GetIPInfoBatch(
	ips []net.IP,
	opts BatchReqOpts,
) (BatchCore, error) {
	return DefaultClient.GetIPInfoBatch(ips, opts)
}

// GetIPInfoBatch does a batch request for all `ips` at once.
func (c *Client) GetIPInfoBatch(
	ips []net.IP,
	opts BatchReqOpts,
) (BatchCore, error) {
	// TODO
	// wrapper over c.GetBatch; convert `ips` to string array, then pass it in,
	// and create a new map which is BatchCore.
	return nil, nil
}

/* ASN */

// GetASNDetailsBatch does a batch request for all `asns` at once.
func GetASNDetailsBatch(
	asns []string,
	opts BatchReqOpts,
) (BatchASNDetails, error) {
	return DefaultClient.GetASNDetailsBatch(asns, opts)
}

// GetASNDetailsBatch does a batch request for all `asns` at once.
func (c *Client) GetASNDetailsBatch(
	asns []string,
	opts BatchReqOpts,
) (BatchASNDetails, error) {
	// TODO
	// wrapper over c.GetBatch; check that `asns` are all ASNs, then pass it
	// in, and create a new map which is BatchASNDetails.
	return nil, nil
}
