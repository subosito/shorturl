package shorturl

import (
	"errors"
	"os"

	"github.com/subosito/shorturl/bitly"
	"github.com/subosito/shorturl/gggg"
	"github.com/subosito/shorturl/gitio"
	"github.com/subosito/shorturl/isgd"
	"github.com/subosito/shorturl/pendekin"
	"github.com/subosito/shorturl/shorl"
	"github.com/subosito/shorturl/tinyurl"
)

type Client struct {
	Provider string
	Params   map[string]string
}

func NewClient(provider string) *Client {
	return &Client{Provider: provider}
}

func (c *Client) Shorten(u string) (string, error) {
	switch c.Provider {
	case "tinyurl":
		s := tinyurl.New()
		return s.Shorten(u)
	case "isgd":
		s := isgd.New()
		return s.Shorten(u)
	case "gitio":
		s := gitio.New()
		return s.Shorten(u)
	case "bitly":
		token := os.Getenv("BITLY_ACCESS_TOKEN")
		s := bitly.New(token)
		return s.Shorten(u)
	case "shorl":
		s := shorl.New()
		return s.Shorten(u)
	case "gggg":
		s := gggg.New()
		return s.Shorten(u)
	case "pendekin":
		s := pendekin.New()
		return s.Shorten(u)
	}

	err := errors.New("You should not see this :P")
	return "", err
}
