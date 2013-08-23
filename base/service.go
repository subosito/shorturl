// package base provides HTTP service implementation to communicates with remote providers.
package base

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Servicer interface {
	Shorten(u string) ([]byte, error)
}

type Service struct {
	Scheme    string
	Host      string
	Code      int
	Method    string
	Path      string
	Field     string
	Transport http.RoundTripper
}

func (s *Service) transport() http.RoundTripper {
	if s.Transport != nil {
		return s.Transport
	}

	return http.DefaultTransport
}

func (s *Service) Url() url.URL {
	return url.URL{
		Scheme: s.Scheme,
		Host:   s.Host,
		Path:   s.Path,
	}
}

func (s *Service) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", err)
	fmt.Printf("%s\n", b)

	return b, nil
}

func (s *Service) Request(u string) (*http.Response, error) {
	su := s.Url()
	d := ""
	v := url.Values{}
	v.Set(s.Field, u)

	switch s.Method {
	case "GET":
		su.RawQuery = v.Encode()
	case "POST":
		d = v.Encode()
	}

	req, err := http.NewRequest(s.Method, su.String(), strings.NewReader(d))

	fmt.Printf("%+v\n", req)

	if err != nil {
		return nil, err
	}

	client := &http.Client{Transport: s.transport()}
	return client.Do(req)
}

func (s *Service) Read(r *http.Response) ([]byte, error) {
	if r.StatusCode != s.Code {
		m := fmt.Sprintf("Unexpected StatusCode: %d != %d", r.StatusCode, s.Code)
		return nil, errors.New(m)
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
