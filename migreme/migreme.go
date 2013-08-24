// package migreme provides support for migre.me shortening service.
package migreme

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Migreme struct {
	*base.Service
}

func New() *Migreme {
	return &Migreme{&base.Service{
		Scheme: "http",
		Host:   "migre.me",
		Method: "GET",
		Path:   "/api.txt",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}
