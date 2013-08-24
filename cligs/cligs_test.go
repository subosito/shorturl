package cligs

import (
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestCligs(t *testing.T) {
	s := New()
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
