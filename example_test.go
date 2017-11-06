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
