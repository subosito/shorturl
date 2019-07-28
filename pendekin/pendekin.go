// package pendekin provides support for pendek.in shortening service.
package pendekin

import (
	"net/http"
	"regexp"

	"github.com/subosito/shorturl/base"
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

func (s *Pendekin) Shorten(u string) (string, error) {
	res, err := s.Request(u)
	if err != nil {
		return "", err
	}

	b, err := s.Read(res)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(Pattern)
	o := r.FindString(string(b))

	return o, nil
}
