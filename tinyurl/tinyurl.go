// package tinyurl provides support for tinyurl.com shortening service.
package tinyurl

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type TinyURL struct {
	*base.Service
}

func New() *TinyURL {
	return &TinyURL{&base.Service{
		Scheme: "http",
		Host:   "tinyurl.com",
		Method: "GET",
		Path:   "/api-create.php",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}
