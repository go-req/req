package req

import "io"

// Response represents the response from an HTTP request.s
type Response struct {
	// Body of the response, closed automatically if Stream is false
	Body io.ReadCloser
}
