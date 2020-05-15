Simple helper to obfuscate serial, integer numbers (id in particular).

[![GoDoc](https://godoc.org/github.com/raohwork/intobfus?status.svg)](https://godoc.org/github.com/raohwork/intobfus)
[![Build Status](https://travis-ci.org/raohwork/intobfus.svg?branch=master)](https://travis-ci.org/raohwork/intobfus)
[![Go Report Card](https://goreportcard.com/badge/github.com/raohwork/intobfus)](https://goreportcard.com/report/github.com/raohwork/intobfus)

# Synopsis

See [examples](https://godoc.org/github.com/raohwork/intobfus/intobfus#pkg-examples)

# CLI helper

```sh
# download binary
go get github.com/raohwork/intobfus

# generate go code template for 1~(2^32-1)
intobfus -pass 3 -bits 32
# or for 1~283476
intobfus -pass 3 -max 283476
# show some data to stderr
intobfus -pass 3 -bits 64 -show
```

# License

MPL v2.0
