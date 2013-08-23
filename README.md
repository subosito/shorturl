# ShortURL

[![Build Status](https://travis-ci.org/subosito/shorturl.png)](https://travis-ci.org/subosito/shorturl)

Generic implementation for interacting with various URL shortening services in Go.

## Usage

There are two ways using shorturl package:

1. Unified client

    ```go
    // import "github.com/subosito/shorturl"

    u := "http://subosito.com/"

    c := shorturl.NewClient("tinyurl") // you can use another provider, eg: "isgd"
    u, err := c.Shorten(Url)
    if err == nil {
    	fmt.Println(u)
    }
    ```

2. Provider

    ```go
    // import "github.com/subosito/shorturl/tinyurl"

    u := "http://subosito.com/"
    s := tinyurl.New()
    u, err := s.Shorten(Url)
    if err == nil {
    	fmt.Println(u)
    }
    ```

## Credits

Inspired by @robbyrussell's shorturl implementation in Ruby, https://github.com/robbyrussell/shorturl.
