// package bitly provides support for bit.ly shortening service.
package bitly

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Bitly struct {
	*base.Service
}

func New() *Bitly {
	return &Bitly{&base.Service{
		Scheme: "http",
		Host:   "api.bitly.com",
		Method: "GET",
		Path:   "/v3/shorten/",
		Field:  "longUrl",
		Code:   http.StatusOK,
		Params: map[string]string{
			"format": "txt",
			"login":  "",
			"apiKey": "",
		},
	}}
}
