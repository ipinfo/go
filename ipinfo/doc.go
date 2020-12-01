/*
Package ipinfo provides a client for using the IPinfo API.

Usage:

	import "github.com/ipinfo/go/ipinfo"

The default IPinfo client is predefined and can be used without initialization.
For example:

	info, err := ipinfo.GetInfo(net.ParseIP("8.8.8.8"))

Authentication

To perform authenticated API calls construct a new IPinfo client using
AuthTransport HTTP client. For example:

	authTransport := ipinfo.AuthTransport{Token: "MY_TOKEN"}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)
	info, err := client.GetInfo(net.ParseIP("8.8.8.8"))

Note that when using an authenticated Client, all calls made by the client will
include the specified token.
*/
package ipinfo
