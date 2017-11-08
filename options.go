package req

// A set of convenience options to use for requests.

// JSON will automatically marshal the field, and set the Content-Type header
// to application/json
func JSON(v interface{}) func(r *Request) {
	return func(r *Request) {
		r.JSON = v
		r.Headers["Content-Type"] = "application/json"
	}
}

// Redirect will set the max amount of times a request is allowed to redirect.
func Redirect(num int) func(r *Request) {
	return func(r *Request) {
		r.NumRedirects = num
	}
}
