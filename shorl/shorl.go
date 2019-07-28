// package shorl provides support for shorl.com shortening service.
package shorl

import (
	"net/http"
	"regexp"

	"github.com/subosito/shorturl/base"
)

var Pattern string = "http://shorl.com/[a-zA-Z0-9]+"

type Shorl struct {
	*base.Service
}

func New() *Shorl {
	return &Shorl{&base.Service{
		Scheme: "http",
		Host:   "shorl.com",
		Method: "GET",
		Path:   "/create.php",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}

func (s *Shorl) Shorten(u string) (string, error) {
	res, err := s.Request(u)
	if err != nil {
		return "", err
	}

	b, err := s.Read(res)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(Pattern)
	m := r.FindString(string(b))
	return m, nil
}
