Fake
====

Fake is a fake data generator for Go (Golang), heavily inspired by the forgery and ffaker Ruby gems.

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/icrowley/fake) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/icrowley/fake/master/LICENSE)

# About

Most data and methods are ported from forgery/ffaker Ruby gems.
For the list of available methods please look at https://godoc.org/github.com/icrowley/fake.
Currently english and russian languages are available.
Feel free to add other languages, but don't forget to regenerate data.go file using `github.com/mjibson/esc` tool and `esc -o data.go -pkg fake data` command because Fake embeds files unless you call `UseExternalData(true)` in order to be able to work without external files dependencies when compiled.

## Install

```shell
go get github.com/icrowley/fake
```

## Import

```go
import (
  "github.com/icrowley/fake"
)
```

## Documentation

Documentation can be found at godoc:

https://godoc.org/github.com/icrowley/fake

## Test
To run the project tests:

```shell
cd test
go test
```

## Examples

```go
name := fake.FirstName()
fullname = := fake.FullName()
product := fake.Product()
```

Changing language:

```go
err := fake.SetLang("ru")
if err != nil {
  panic(err)
}
password := fake.SimplePassword()
```

Using english fallback:

```go
err := fake.SetLang("ru")
if err != nil {
  panic(err)
}
fake.EnFallback(true)
password := fake.Paragraph()
```

Using external data:

```go
fake.UseExternalData(true)
password := fake.Paragraph()
```

## Author

Dmitry Afanasyev,
http://twitter.com/i_crowley
dimarzio1986@gmail.com
