package ipinfo

import (
	"net"
	"sync"
)

const (
	batchMaxSize           = 1000
	batchReqTimeoutDefault = 5
)

// Batch is a mapped result of any valid API endpoint (e.g. `<ip>`,
// `<ip>/<field>`, `<asn>`, etc) to its corresponding data.
type Batch map[string]interface{}

// BatchCore is a mapped result of IPs to their corresponding `Core` data.
type BatchCore map[net.IP]*Core

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
	var batchSize

	// use correct batch size; default/clip to `batchMaxSize`.
	if opts.BatchSize == 0 || opts.BatchSize > batchMaxSize {
		batchSize = batchMaxSize
	} else {
		batchSize = opts.BatchSize
	}
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
}
