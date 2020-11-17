package ipinfo // import "github.com/ipinfo/go/ipinfo"

import (
	"net/http"
)

// AuthTransport is an http.RoundTripper that authenticates all requests with a
// token query string.
type AuthTransport struct {
	Token     string
	Transport http.RoundTripper
}

// RoundTrip implements the RoundTripper interface.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req = cloneRequest(req)
	q := req.URL.Query()
	q.Set("token", t.Token)
	req.URL.RawQuery = q.Encode()
	return t.transport().RoundTrip(req)
}

// Client returns an *http.Client that makes requests that are authenticated
// with a token query string.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func cloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
