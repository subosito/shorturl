# ShortURL

[![Build Status](https://travis-ci.org/subosito/shorturl.svg?branch=master)](https://travis-ci.org/subosito/shorturl)
[![GoDoc](https://godoc.org/github.com/subosito/shorturl?status.svg)](https://godoc.org/github.com/subosito/shorturl)

Generic implementation for interacting with various URL shortening services in Go.

## Usage

As usual you can install the package by issuing:

```bash
$ go get github.com/subosito/shorturl
```

## Features

**Shortening URL**

ShortURL provides simple API to shorten a long URL, here's how:

```go
// import "github.com/subosito/shorturl"

provider := "tinyurl"
u, err := shorturl.Shorten("http://example.com/", provider)
if err == nil {
	fmt.Println(u)
}
```

Alternatively, you can initialize desired provider:

```go
// import "github.com/subosito/shorturl/tinyurl"

s := tinyurl.New()
u, err := s.Shorten("http://example.com/")
if err == nil {
	fmt.Println(u)
}
```

**Expanding short URL**

ShortURL also provides simple API to expand short URL into its original long URL:

```go
// import "github.com/subosito/shorturl"

u, err := shorturl.Expand("http://bit.ly/13M3JX5")
if err == nil {
	fmt.Println(u)
}
```

## Supported Services

| Package     | Service                  | Enviroment Variables       |
|-------------|--------------------------|----------------------------|
| `bitly`     | https://bitly.com/       | `BITLY_ACCESS_TOKEN`       |
| `gggg`      | http://gg.gg/            | -                          |
| `gitio`     | https://git.io/          | -                          |
| `isgd`      | https://is.gd/           | -                          |
| `shorl`     | http://shorl.com/        | -                          |
| `tinyurl`   | https://tinyurl.com/     | -                          |

## Credits

Inspiration comes from Rubygem [shorturl](https://github.com/robbyrussell/shorturl) by [@robbyrussell](https://github.com/robbyrussell)

