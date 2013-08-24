// package parapt provides support for para.pt shortening service.
package parapt

import (
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Parapt struct {
	*base.Service
}

func New() *Parapt {
	return &Parapt{&base.Service{
		Scheme: "http",
		Host:   "para.pt",
		Method: "GET",
		Path:   "/api.aspx",
		Field:  "url",
		Params: map[string]string{
			"alias": "",
			"cmd":   "CREATE",
		},
		Code: http.StatusOK,
	}}
}
