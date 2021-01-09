package ipinfo

import (
	"net"
)

// Batch is a mapped result of IPs to their corresponding `Core` data.
type Batch map[net.IP]*Core

// BatchReqOpts are options input into the `GetIPInfoBatch` functions.
type BatchReqOpts {
	// BatchSize is the internal batch size used per API request; the IPinfo
	// API has a maximum batch size, but the GetIPInfoBatch functions available
	// in this library do not. Therefore the library chunks the IP array
	// internally into chunks of size `BatchSize`, clipping to the maximum
	// allowed by the IPinfo API.
	BatchSize uint32

	// TimeoutPerBatch is the timeout in seconds that each batch of size
	// `BatchSize` will have for its own request.
	//
	// 0 means no timeout at all per batch request; this does _not_ override
	// the value of `TimeoutTotal` if that is non-0.
	TimeoutPerBatch uint64

	// TimeoutTotal is the total timeout in seconds for all batch requests in a
	// `GetIPInfoBatch` function to complete.
	//
	// 0 means no total timeout; `TimeoutPerBatch` will still apply.
	TimeoutTotal uint64
}

// GetIPInfoBatch does a batch request for all `ips` at once.
func GetIPInfoBatch(
	ips []net.IP,
	opts BatchReqOpts,
) (Batch, error) {
	return DefaultClient.GetIPInfoBatch(ips, opts)
}

// GetIPInfoBatch does a batch request for all `ips` at once.
func (c *Client) GetIPInfoBatch(
	ips []net.IP,
	opts BatchReqOpts,
) (Batch, error) {

}
