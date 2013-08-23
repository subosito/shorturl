package shorturl

import (
	"errors"
	"github.com/subosito/shorturl/isgd"
)

type Client struct {
	Provider string
	Params   map[string]string
}

func NewClient(provider string) *Client {
	return &Client{Provider: provider}
}

func (c *Client) Shorten(u string) ([]byte, error) {
	switch c.Provider {
	case "isgd":
		s := isgd.New()
		return s.Shorten(u)
	}

	err := errors.New("You should not see this :P")
	return nil, err
}
