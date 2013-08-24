// package gggg provides support for gg.gg shortening service.
package gggg

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Gggg struct {
	*base.Service
}

func New() *Gggg {
	return &Gggg{&base.Service{
		Scheme: "http",
		Host:   "gg.gg",
		Method: "POST",
		Path:   "/create",
		Field:  "long_url",
		Code:   http.StatusOK,
		Params: map[string]string{
			"app":         "site",
			"version":     "0.1",
			"custom_path": "",
		},
	}}
}
