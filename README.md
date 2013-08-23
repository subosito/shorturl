# ShortURL

[![Build Status](https://travis-ci.org/subosito/shorturl.png)](https://travis-ci.org/subosito/shorturl)

Generic implementation for interacting with various URL shortening services in Go.

## Usage

There are two ways using shorturl package:

#### Unified client

```go
// import "github.com/subosito/shorturl"

u := "http://subosito.com/"

c := shorturl.NewClient("tinyurl") // you can use another provider, eg: "isgd"
u, err := c.Shorten(Url)
if err == nil {
fmt.Println(u)
}
```

#### Provider

```go
// import "github.com/subosito/shorturl/tinyurl"

u := "http://subosito.com/"
s := tinyurl.New()
u, err := s.Shorten(Url)
if err == nil {
fmt.Println(u)
}
```

## Supported Services

- [tinyurl](http://tinyurl.com/)
- [is.gd](http://is.gd/)
- [git.io](http://git.io/)
- [bit.ly](https://bitly.com/)
- [ln-s](http://ln-s.net/)
- [shorl](http://shorl.com/)
- [va.mu](http://va.mu/)
- [moourl](http://moourl.com/)
- ...

## Resources

- [Documentation](http://godoc.org/github.com/subosito/shorturl)

## Credits

Inspiration comes from [shorturl Rubygem](https://github.com/robbyrussell/shorturl) by [@robbyrussell](https://github.com/robbyrussell)

