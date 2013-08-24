// package pendekin provides support for pendek.in shortening service.
package pendekin

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"regexp"
)

var Pattern string = `http://pendek.in/[a-z0-9A-Z]+`

type Pendekin struct {
	*base.Service
}

func New() *Pendekin {
	return &Pendekin{&base.Service{
		Scheme: "http",
		Host:   "pendek.in",
		Method: "GET",
		Path:   "/index.php",
		Field:  "longurl",
		Code:   http.StatusOK,
	}}
}

func (s *Pendekin) Shorten(u string) ([]byte, error) {
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

	return []byte(o), nil
}
