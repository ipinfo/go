# go-ipinfo

go-ipinfo is a Go client library for accessing the [IPinfo API](https://ipinfo.io/).

## Usage

```go
import "github.com/ipinfo/go-ipinfo/ipinfo"
```

The default IPinfo client is predefined and can be used without initialization.
For example:


```go
info, err := ipinfo.GetInfo(net.ParseIP("8.8.8.8"))
```

### Authentication

To perform authenticated API calls construct a new IPinfo client using
AuthTransport HTTP client. For example:


```go
authTransport := ipinfo.AuthTransport{Token: "MY_TOKEN"}
httpClient := authTransport.Client()
client := ipinfo.NewClient(httpClient)
info, err := client.GetInfo(net.ParseIP("8.8.8.8"))
```

Note that when using an authenticated Client, all calls made by the client will
include the specified token.
