package ipinfo

import (
	"net"
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
	// in this library do not. Therefore the library chunks the IP array
	// internally into chunks of size `BatchSize`, clipping to the maximum
	// allowed by the IPinfo API.
	//
	// 0 means to use the default batch size which is the max allowed by the
	// IPinfo API.
	BatchSize uint32

	// TimeoutPerBatch is the timeout in seconds that each batch of size
	// `BatchSize` will have for its own request.
	//
	// 0 means no timeout at all per batch request; this does _not_ override
	// the value of `TimeoutTotal` if that is non-0.
	TimeoutPerBatch uint64

	// TimeoutTotal is the total timeout in seconds for all batch requests in a
	// batch request function to complete.
	//
	// 0 means no total timeout; `TimeoutPerBatch` will still apply.
	TimeoutTotal uint64

	// Filter, if turned on, will filter out a URL whose value was deemed empty
	// on the server.
	Filter bool
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

}
