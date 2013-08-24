package snipurl

import (
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestSnipURL(t *testing.T) {
	s := New()
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
