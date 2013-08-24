// package vamu provides support for va.mu shortening service.
package vamu

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Vamu struct {
	*base.Service
}

func New() *Vamu {
	return &Vamu{&base.Service{
		Scheme: "http",
		Host:   "va.mu",
		Method: "GET",
		Path:   "/api/create",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}
