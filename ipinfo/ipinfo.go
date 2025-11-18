package ipinfo

// DefaultClient is the package-level client available to the user.
var DefaultClient *Client

// Package-level Lite bundle client
var DefaultLiteClient *LiteClient

// Package-level Core bundle client
var DefaultCoreClient *CoreClient

func init() {
	// Create global clients for legacy, Lite, Core APIs
	DefaultClient = NewClient(nil, nil, "")
	DefaultLiteClient = NewLiteClient(nil, nil, "")
	DefaultCoreClient = NewCoreClient(nil, nil, "")
}
