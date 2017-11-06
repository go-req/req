package req

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
	"net/url"
)

// go-req version
const version = "1.0"

// Client is the default client to use for requests
var Client = &http.Client{}

// Headers is an alias to net/http's Header
type Headers = http.Header

// Params is an alias to net/url's Values
type Params = url.Values

// A Request represents an HTTP request to be sent.
// See NewRequest to view the defaults for a Request
type Request struct {
	// The HTTP Method
	Method string
	// URL to send the request to
	URL string

	// Map of HTTP Headers to send with the request
	Headers Headers
	// Map of url params to modify the url
	Params Params

	// Leave the body open, and don't automatically read the text.
	Stream bool
	// HTTP Timeout, defaults to 60 seconds
	Timeout time.Duration
	// Number of redirects to allow, defaults to 10.
	NumRedirects int

	// Transport to use for the client.
	Transport *http.Transport

	// Body is the request's body, this takes precedence over Content
	Body io.Reader
	// Content of the HTTP Request, this takes precedence over Text
	Content []byte
	// Text to send in the HTTP Request, this takes precedence over JSON
	Text string
	// Req will automatically marshal this into JSON
	JSON interface{}
}

// Do will execute the request.
// Creates a copy of the client to avoid modifying it directly.
func (req *Request) Do() (*Response, error) {
	var body io.Reader

	/*
		Order of precendence for the request body
			1: Body
			2: Content
			3: Text
			4: JSON
	*/
	switch {
	case req.Body != nil:
		body = req.Body
	case req.Content != nil:
		body = bytes.NewBuffer(req.Content)
	case req.Text != "":
		body = bytes.NewBufferString(req.Text)
	case req.JSON:
		b, err := json.Marshal(req.JSON)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(b)
	default:
		body = nil
	}

	r, err := http.NewRequest(req.Method, req.URL, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	*client = *Client

	if req.Transport != nil {
		client.Transport = req.Transport
	}

	client.Timeout = req.Timeout

	// Stop redirects after NumRedirects
	client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		if len(via) >= req.NumRedirects {
			return http.ErrUseLastResponse
		}
		return nil
	}

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	return toResponse(req, resp)
}

// NewRequest will return a Request with the preferred defaults
// without executing it.
func NewRequest() *Request {
	return &Request{
		Headers: http.Header{
			"Accept":     []string{"*/*"},
			"User-Agent": []string{"go-req/" + version},
		},
		NumRedirects: 10,
		Timeout: 60 * time.Second,
	}
}

// New will construct a request and immediately execute it
func New(method, url string, options ...func(*Request)) (*Response, error) {
	req := NewRequest()

	req.Method = method
	req.URL = url

	for _, fn := range options {
		fn(req)
	}

	return req.Do()
}

// Get will send a GET request
func Get(url string, options ...func(*Request)) (*Response, error) {
	return New("GET", url, options...)
}

// Post will send a POST request
func Post(url string, options ...func(*Request)) (*Response, error) {
	return New("POST", url, options...)
}

// Put will send a PUT request
func Put(url string, options ...func(*Request)) (*Response, error) {
	return New("PUT", url, options...)
}

// Patch will send a PATCH request
func Patch(url string, options ...func(*Request)) (*Response, error) {
	return New("PATCH", url, options...)
}

// Delete will send a DELETE request
func Delete(url string, options ...func(*Request)) (*Response, error) {
	return New("DELETE", url, options...)
}
