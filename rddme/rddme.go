// package rdd.me provides support for rdd.me, a readability's shortening service.
package rddme

import (
	"encoding/json"
	"github.com/subosito/shorturl/base"
	"net/http"
)

type Rddme struct {
	*base.Service
}

type metaResponse struct {
	URL    string `json:"url"`
	RddURL string `json:"rdd_url"`
	ID     string `json:"id"`
}

type response struct {
	Meta     metaResponse `json:"meta"`
	Messages []string     `json:"messages"`
	Success  bool         `json:"success"`
}

func New() *Rddme {
	return &Rddme{&base.Service{
		Scheme: "http",
		Host:   "www.readability.com",
		Method: "POST",
		Path:   "/api/shortener/v1/urls",
		Field:  "url",
		Code:   http.StatusAccepted,
	}}
}

func (s *Rddme) Shorten(u string) ([]byte, error) {
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

	return []byte(r.Meta.RddURL), nil
}
