package req

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// Response represents the response from an HTTP request.s
type Response struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Ok         bool   // If the StatusCode is >= 200 and < 400

	// The Raw response provided by net/http
	Raw *http.Response

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	Request *Request

	// Bytes populated from the response body.
	// This is nil if Stream was set to true.
	Content []byte
	// Response data as a string
	Text string

	// Body of the response, closed automatically if Stream is false
	Body io.ReadCloser
	// Headers the server responded with
	Headers http.Header
	// A Cookie represents an HTTP cookie as sent in the Set-Cookie header in the
	// HTTP response.
	Cookies []*http.Cookie

	// Uncompressed reports whether the response was sent compressed but
	// was decompressed by the http package. When true, reading from
	// Body yields the uncompressed content instead of the compressed
	// content actually set from the server, ContentLength is set to -1,
	// and the "Content-Length" and "Content-Encoding" fields are deleted
	// from the responseHeader. To get the original response from
	// the server, set Client.Transport.DisableCompression to true.
	Uncompressed bool
}

// JSON is shorthand for json.Unmarshal(r.Content, to)
func (r *Response) JSON(to interface{}) error {
	return json.Unmarshal(r.Content, to)
}

// Convert a net/http response to a req.Response.
// Body is closed before any error can occur if Stream is false.
// Only error that can occur is when reading from the body fails.
func toResponse(req *Request, resp *http.Response) (*Response, error) {
	var bytes []byte
	if !req.Stream {
		// Default
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		bytes = b
	}

	var ok bool
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		ok = true
	}

	return &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Ok:         ok,

		Raw:     resp,
		Request: req,

		Content: bytes,
		Text:    string(bytes),
		Body:    resp.Body,
		Headers: resp.Header,
		Cookies: resp.Cookies(),
	}, nil
}
