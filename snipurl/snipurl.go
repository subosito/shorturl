// package snipurl provides support for snipurl.com shortening service.
package snipurl

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"regexp"
	"strings"
)

var Pattern string = `value="http://snipurl.com/[0-9a-zA-Z]+`

type SnipURL struct {
	*base.Service
}

func New() *SnipURL {
	return &SnipURL{&base.Service{
		Scheme: "http",
		Host:   "snipurl.com",
		Method: "POST",
		Field:  "url",
		Code:   http.StatusOK,
		Params: map[string]string{
			"alias":       "",
			"title":       "",
			"private_key": "",
		},
	}}
}

func (s *SnipURL) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(Pattern)
	o := r.FindString(string(b))
	f := strings.Replace(o, `value="`, "", 1)

	return []byte(f), nil
}
