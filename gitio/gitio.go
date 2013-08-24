// package gitio provides support for git.io, a Github URL shortening services.
// You can shorten https://github.com/* and https://gist.github.com/* url only.
package gitio

import (
	"github.com/subosito/shorturl/base"
	"net/http"
	"net/url"
)

type GitIO struct {
	*base.Service
}

func New() *GitIO {
	return &GitIO{&base.Service{
		Scheme: "http",
		Host:   "git.io",
		Method: "POST",
		Path:   "/create",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}

func (s *GitIO) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	ux := url.URL{
		Scheme: s.Scheme,
		Host:   s.Host,
		Path:   string(append([]byte("/"), b...)),
	}

	return []byte(ux.String()), nil
}
