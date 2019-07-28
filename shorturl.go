package shorturl

import (
	"net/http"
)

func Shorten(u, provider string) (string, error) {
	c := NewClient(provider)
	return c.Shorten(u)
}

func Expand(u string) (string, error) {
	res, err := http.Get(u)
	if err != nil {
		return "", err
	}

	f := res.Request.URL.String()
	return f, nil
}
