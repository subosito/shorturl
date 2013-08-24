// package moourl provides support for moourl.com shortening service.
package moourl

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"strings"
)

type MooURL struct {
	*base.Service
}

func New() *MooURL {
	return &MooURL{&base.Service{
		Scheme: "http",
		Host:   "moourl.com",
		Method: "GET",
		Path:   "/create/",
		Field:  "source",
		Code:   http.StatusFound,
	}}
}

func (s *MooURL) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	x := res.Request.URL.String()
	f := strings.Replace(x, "woot/?moo=", "", 1)
	return []byte(f), nil
}
