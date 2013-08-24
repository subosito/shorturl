// package catchy provides support for catchylink.com shortening service.
package catchy

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"regexp"
	"strings"
)

var Pattern string = `value="http://catchylink.com/[a-z0-9A-Z]+`

type Catchy struct {
	*base.Service
}

func New() *Catchy {
	return &Catchy{&base.Service{
		Scheme: "http",
		Host:   "www.catchylink.com",
		Method: "GET",
		Path:   "/create",
		Field:  "u",
		Code:   http.StatusOK,
	}}
}

func (s *Catchy) Shorten(u string) ([]byte, error) {
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
