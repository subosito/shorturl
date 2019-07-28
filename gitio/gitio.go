// package gitio provides support for git.io, a Github URL shortening services.
// You can shorten https://github.com/* and https://gist.github.com/* url only.
package gitio

import (
	"net/http"
	"net/url"

	"github.com/subosito/shorturl/base"
)

type GitIO struct {
	*base.Service
}

func New() *GitIO {
	return &GitIO{&base.Service{
		Scheme: "https",
		Host:   "git.io",
		Method: "POST",
		Path:   "/create",
		Field:  "url",
		Code:   http.StatusOK,
	}}
}

func (s *GitIO) Shorten(u string) (string, error) {
	res, err := s.Request(u)
	if err != nil {
		return "", err
	}

	b, err := s.Read(res)
	if err != nil {
		return "", err
	}

	ux := url.URL{
		Scheme: s.Scheme,
		Host:   s.Host,
		Path:   string(append([]byte("/"), b...)),
	}

	return ux.String(), nil
}
