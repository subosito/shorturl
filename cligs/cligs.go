// package cligs provides support for cli.gs shortening service.
package cligs

import (
	"github.com/subosito/shorturl/base"
	"regexp"
	"strings"
)

type Cligs struct {
	*base.Service
}

func New() *Cligs {
	return &Cligs{&base.Service{
		Scheme: "http",
		Host:   "cli.gs",
		Method: "POST",
		Field:  "url",
		Code:   200,
		Params: map[string]string{
			"do":         "shorten",
			"sbmt":       "Shorten",
			"custom_url": "",
		},
	}}
}

func (s *Cligs) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`value="http://cli.gs/[0-9a-zA-Z]+`)
	o := r.FindString(string(b))
	f := strings.Replace(o, `value="`, "", 1)

	return []byte(f), nil
}
