package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/snail007/gf/net/ghttp"
)

func main() {
	c := ghttp.NewClient()
	c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	r, e := c.Clone().Get("https://127.0.0.1:8199")
	fmt.Println(e)
	fmt.Println(r.StatusCode)
}
