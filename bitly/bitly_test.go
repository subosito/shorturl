package bitly

import (
	"os"
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestBitly(t *testing.T) {
	s := New()
	s.Params["login"] = os.Getenv("BITLY_LOGIN")
	s.Params["apiKey"] = os.Getenv("BITLY_API_KEY")
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
