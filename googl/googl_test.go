package googl

import (
	"os"
	"testing"
)

var Url string = "http://github.com/subosito/shorturl"

func TestGoogl(t *testing.T) {
	s := New()
	s.Params["key"] = os.Getenv("GOOGL_API_KEY")
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
