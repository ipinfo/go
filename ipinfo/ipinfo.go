package ipinfo

// DefaultClient is the package-level client available to the user.
var DefaultClient *Client

// Package-level lite client
var DefaultLiteClient *LiteClient

func init() {
	// Create two global clients, one for Core and one for Lite API
	DefaultClient = NewClient(nil, nil, "")
	DefaultLiteClient = NewLiteClient(nil, nil, "")
}
