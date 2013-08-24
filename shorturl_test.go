package shorturl

import (
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"
var ShortUrl string = "http://bit.ly/13M3JX5"

func TestShortURL_Shorten(t *testing.T) {
	u, err := Shorten(Url, "tinyurl")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}

func TestShortURL_Expand(t *testing.T) {
	u, err := Expand(ShortUrl)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", ShortUrl, u)
}
