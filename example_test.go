package req_test

import (
	"fmt"

	"github.com/go-req/req"
)

func ExampleGet() {
	// Don't forget to check errors!
	resp, _ := req.Get("https://httpbin.org/get")

	var m map[string]string
	resp.JSON(&m)

	fmt.Println(m["url"])
	// Output: https://httpbin.org/get
}

func ExampleRedirect() {
	// Don't forget to check errors!

	resp, _ := req.Get("https://httpbin.org/redirect/3", req.Redirect(2))

	fmt.Println(resp.StatusCode) // 302 status code means we ended on a request trying to redirect us
	// Output: 302
}
