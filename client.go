package shorturl

import (
	"errors"
	"github.com/subosito/shorturl/adfly"
	"github.com/subosito/shorturl/bitly"
	"github.com/subosito/shorturl/catchy"
	"github.com/subosito/shorturl/cligs"
	"github.com/subosito/shorturl/gggg"
	"github.com/subosito/shorturl/gitio"
	"github.com/subosito/shorturl/googl"
	"github.com/subosito/shorturl/isgd"
	"github.com/subosito/shorturl/moourl"
	"github.com/subosito/shorturl/pendekin"
	"github.com/subosito/shorturl/shorl"
	"github.com/subosito/shorturl/snipurl"
	"github.com/subosito/shorturl/tinyurl"
	"github.com/subosito/shorturl/vamu"
	"os"
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
		s := bitly.New()
		s.Params["login"] = os.Getenv("BITLY_LOGIN")
		s.Params["apiKey"] = os.Getenv("BITLY_API_KEY")
		return s.Shorten(u)
	case "shorl":
		s := shorl.New()
		return s.Shorten(u)
	case "vamu":
		s := vamu.New()
		return s.Shorten(u)
	case "moourl":
		s := moourl.New()
		return s.Shorten(u)
	case "cligs":
		s := cligs.New()
		return s.Shorten(u)
	case "snipurl":
		s := snipurl.New()
		return s.Shorten(u)
	case "adfly":
		s := adfly.New()
		return s.Shorten(u)
	case "googl":
		s := googl.New()
		s.Params["key"] = os.Getenv("GOOGL_API_KEY")
		return s.Shorten(u)
	case "gggg":
		s := gggg.New()
		return s.Shorten(u)
	case "pendekin":
		s := pendekin.New()
		return s.Shorten(u)
	case "catchy":
		s := catchy.New()
		return s.Shorten(u)
	}

	err := errors.New("You should not see this :P")
	return nil, err
}
