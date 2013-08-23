package shorturl

import (
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestShortURL(t *testing.T) {
	c := NewClient("isgd")
	u, err := c.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
