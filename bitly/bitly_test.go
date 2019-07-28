package bitly

import (
	"os"
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestBitly(t *testing.T) {
	token := os.Getenv("BITLY_ACCESS_TOKEN")

	s := New(token)
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
