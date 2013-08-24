package shorturl

import (
	"net/http"
)

func Shorten(u, provider string) ([]byte, error) {
	c := NewClient(provider)
	return c.Shorten(u)
}

func Expand(u string) ([]byte, error) {
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	f := res.Request.URL.String()
	return []byte(f), nil
}
