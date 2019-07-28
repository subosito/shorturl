// package bitly provides support for bit.ly shortening service.
package bitly

import (
	"encoding/json"
	"net/http"

	"github.com/subosito/shorturl/base"
)

type Bitly struct {
	*base.Service
}

func New(token string) *Bitly {
	return &Bitly{&base.Service{
		Scheme: "https",
		Host:   "api-ssl.bitly.com",
		Method: "POST",
		Format: "JSON",
		Path:   "/v4/shorten",
		Field:  "long_url",
		Code:   http.StatusOK,
		Headers: map[string]string{
			"Authorization": "Bearer " + token,
		},
	}}
}

func (s *Bitly) Shorten(u string) (string, error) {
	res, err := s.Request(u)
	if err != nil {
		return "", err
	}

	b, err := s.Read(res)
	if err != nil {
		return "", err
	}

	v := struct {
		Link string `json:"link"`
	}{}

	if err := json.Unmarshal(b, &v); err != nil {
		return "", err
	}

	return v.Link, nil
}
