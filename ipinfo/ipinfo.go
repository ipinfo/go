//go:generate go run gen-fields.go

package ipinfo

// Global, default client available to the user via `ipinfo.DefaultClient`.
var DefaultClient *Client

func init() {
	/* Create a global, default client. */
	DefaultClient = NewClient(nil, nil, "")
}
