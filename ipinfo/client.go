package ipinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://ipinfo.io/"
	libraryVersion = "1.0"
	userAgent      = "IPinfoClient/Go/" + libraryVersion
)

// A Client manages communication with IPinfo API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. BaseURL should always be specified with a
	// trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the IPinfo API.
	UserAgent string
}

// NewClient returns a new IPinfo API client. If a nil httpClient is provided,
// http.DefaultClient will be used. To use authenticated API methods provide
// http.Client with AuthTransport.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client. Relative
// URLs should always be specified without a preceding slash.
func (c *Client) NewRequest(urlStr string) (*http.Request, error) {
	u := new(url.URL)

	if rel, err := url.Parse(urlStr); err == nil {
		u = c.BaseURL.ResolveReference(rel)
	} else if strings.ContainsRune(urlStr, ':') {
		// IPv6 strings fail to parse as URLs, so let's add it as an
		// URL Path.
		*u = *c.BaseURL
		u.Path += urlStr
	} else {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}

// An ErrorResponse reports an error caused by an API request.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Err      struct {
		Title   string `json:"title"`
		Message string `json:"message"`
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Err)
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

// GetClient get the global Client instance.
func GetClient() *Client {
	return c
}
