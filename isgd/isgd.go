// package isgd provides support for is.gd shortening service.
package isgd

import (
	"github.com/subosito/shorturl/base"
)

type Isgd struct {
	*base.Service
}

func New() *Isgd {
	return &Isgd{&base.Service{
		Scheme: "http",
		Host:   "is.gd",
		Method: "GET",
		Path:   "/api.php",
		Field:  "longurl",
		Code:   200,
	}}
}
