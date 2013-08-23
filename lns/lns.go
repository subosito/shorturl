// package lns provides support for ln-s.net shortening service.
package lns

import (
	"github.com/subosito/shorturl/base"
)

type Lns struct {
	*base.Service
}

func New() *Lns {
	return &Lns {&base.Service{
		Scheme: "http",
		Host:   "ln-s.net",
		Method: "GET",
		Path:   "/home/api.jsp",
		Field:  "url",
		Code:   200,
	}}
}

func (s *Lns) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	return b[4:len(b)], nil
}
