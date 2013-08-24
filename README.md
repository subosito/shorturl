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

| package   | service                |
|-----------|------------------------|
| `tinyurl` | http://tinyurl.com/    |
| `isgd`    | http://is.gd/          |
| `gitio`   | http://git.io/         |
| `bitly`   | https://bitly.com/     |
| `lns`     | http://ln-s.net/       |
| `shorl`   | http://shorl.com/      |
| `vamu`    | http://va.mu/          |
| `moourl`  | http://moourl.com/     |
| `cligs`   | http://cli.gs/         |

## Resources

- [Documentation](http://godoc.org/github.com/subosito/shorturl)

## Credits

Inspiration comes from [shorturl Rubygem](https://github.com/robbyrussell/shorturl) by [@robbyrussell](https://github.com/robbyrussell)

