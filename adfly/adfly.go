// package adfly provides support for adf.ly shortening service.
package adfly

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"regexp"
	"strings"
)

var Pattern string = `http:\\/\\/adf.ly\\/[a-zA-z0-9]+`

type Adfly struct {
	*base.Service
}

func New() *Adfly {
	return &Adfly{&base.Service{
		Scheme: "http",
		Host:   "adf.ly",
		Method: "POST",
		Path:   "/shortener/shorten",
		Field:  "url",
		Code:   http.StatusOK,
		Params: map[string]string{
			"_user_id": "-1",
			"_api_key": "2ba3f6ce601d043c177eb2a83eb34f5g",
		},
	}}
}

func (s *Adfly) Shorten(u string) ([]byte, error) {
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
	f := strings.Replace(o, `\`, ``, -1)

	return []byte(f), nil
}
