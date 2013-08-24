// package googl provides support for goo.gl shortening service.
package googl

import (
	"encoding/json"
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Googl struct {
	*base.Service
}

type response struct {
	Kind    string
	ID      string
	LongURL string
}

func New() *Googl {
	return &Googl{&base.Service{
		Scheme: "https",
		Host:   "www.googleapis.com",
		Method: "POST",
		Path:   "/urlshortener/v1/url",
		Field:  "longUrl",
		Code:   http.StatusOK,
		Format: "JSON",
		Params: map[string]string{
			"key": "",
		},
	}}
}

func (s *Googl) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	r := &response{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}

	return []byte(r.ID), nil
}
