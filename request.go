package req

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// go-req version
const version = "1.0"

// Client is the default client to use for requests
var Client = &http.Client{}

// A Request represents an HTTP request to be sent.
type Request struct {
	// The HTTP Method
	Method string
	// URL to send the request to
	URL string

	// Map of HTTP Headers to send with the request
	Headers http.Header
	// Map of url params to modify the url
	Params url.Values

	// Leave the body open, and don't automatically read the text.
	Stream bool
	// HTTP Timeout
	Timeout int

	// Specific Client to use, overrides the default go-req client
	Client *http.Client

	// Body is the request's body, this takes precedence over Content
	Body io.Reader
	// Content of the HTTP Request, this takes precedence over Text
	Content []byte
	// Text to send in the HTTP Request, this takes precedence over JSON
	Text string
	// Req will automatically marshal this into JSON
	JSON interface{}
}

// Do will execute the request
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

	var client *http.Client

	// If they passed a client, then override client to use.
	if req.Client == nil {
		client = Client
	} else {
		client = req.Client
	}

	resp, err := client.Do(req)

	if err != nil {
		return resp, err
	}
}

// NewRequest will return a Request with the preferred defaults
// without executing it.
func NewRequest() *Request {
	return &Request{
		Headers: http.Header{
			"User-Agent": []string{"go-req/" + version},
		},
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

	return req.Do(req)
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

// Head will send a HEAD request
func Head(url string, options ...func(*Request)) (*Response, error) {
	return New("HEAD", url, options...)
}

// Options will send a OPTIONS request
func Options(url string, options ...func(*Request)) (*Response, error) {
	return New("OPTIONS", url, options...)
}
