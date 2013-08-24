# ShortURL

[![Build Status](https://travis-ci.org/subosito/shorturl.png)](https://travis-ci.org/subosito/shorturl)

Generic implementation for interacting with various URL shortening services in Go.

## Shortening URL

ShortURL provides simple API to shorten a long URL. There are two ways to perform this:

#### Unified client

```go
// import "github.com/subosito/shorturl"

provider := "tinyurl"
u, err := shorturl.Shorten("http://example.com/", provider)
if err == nil {
	fmt.Println(u)
}
```

#### Provider package

```go
// import "github.com/subosito/shorturl/tinyurl"

s := tinyurl.New()
u, err := s.Shorten("http://example.com/")
if err == nil {
	fmt.Println(u)
}
```

## Expanding short URL

Besides shortening long URL, this package also provides simple API to expand short URL into its original long URL.

```go
// import "github.com/subosito/shorturl"

u, err := shorturl.Expand("http://bit.ly/13M3JX5")
if err == nil {
	fmt.Println(u)
}
```

_Notes: Currently, adf.ly is not supported._

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

