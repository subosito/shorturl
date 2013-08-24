# ShortURL

[![Build Status](https://travis-ci.org/subosito/shorturl.png)](https://travis-ci.org/subosito/shorturl)

Generic implementation for interacting with various URL shortening services in Go.

## Usage

There are two ways using shorturl package:

#### Unified client

```go
// import "github.com/subosito/shorturl"

c := shorturl.NewClient("tinyurl") // you can use another provider, eg: "isgd"
u, err := c.Shorten("http://example.com/")
if err == nil {
	fmt.Println(u)
}
```

#### Provider

```go
// import "github.com/subosito/shorturl/tinyurl"

s := tinyurl.New()
u, err := s.Shorten("http://example.com/")
if err == nil {
	fmt.Println(u)
}
```

## Supported Services

| package     | service                | Enviroment Variables           |
|-------------|------------------------|--------------------------------|
| `tinyurl`   | http://tinyurl.com/    | -                              |
| `isgd`      | http://is.gd/          | -                              |
| `gitio`     | http://git.io/         | -                              |
| `bitly`     | https://bitly.com/     | `BITLY_API_KEY`, `BITLY_LOGIN` |
| `lns`       | http://ln-s.net/       | -                              |
| `shorl`     | http://shorl.com/      | -                              |
| `vamu`      | http://va.mu/          | -                              |
| `moourl`    | http://moourl.com/     | -                              |
| `cligs`     | http://cli.gs/         | -                              |
| `snipurl`   | http://snipurl.com/    | -                              |
| `adfly`     | http://adf.ly/         | -                              |
| `migreme`   | http://migre.me/       | -                              |
| `googl`     | http://goo.gl/         | `GOOGL_API_KEY`                |
| `gggg`      | http://gg.gg/          | -                              |

## Resources

- [Documentation](http://godoc.org/github.com/subosito/shorturl)

## Credits

Inspiration comes from [shorturl Rubygem](https://github.com/robbyrussell/shorturl) by [@robbyrussell](https://github.com/robbyrussell)

