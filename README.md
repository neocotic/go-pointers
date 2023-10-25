# go-pointers

[![Static Badge](https://img.shields.io/badge/go.dev-reference-007d9c?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/github.com/neocotic/go-pointers)
[![Build Status](https://img.shields.io/github/actions/workflow/status/neocotic/go-pointers/ci.yml?style=for-the-badge)](https://github.com/neocotic/go-pointers/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/neocotic/go-pointers?style=for-the-badge)](https://github.com/neocotic/go-pointers)
[![License](https://img.shields.io/github/license/neocotic/go-pointers?style=for-the-badge)](https://github.com/neocotic/go-pointers/blob/main/LICENSE.md)

Easy-to-use functions for working with pointers in Go (golang).

## Installation

Install using [go install](https://go.dev/ref/mod#go-install):

``` sh
go install github.com/neocotic/go-pointers
```

Then import the package into your own code:

``` go
import "github.com/neocotic/go-pointers"
```

## Documentation

Documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/neocotic/go-pointers#section-documentation). It
contains an overview and reference.

### Example

``` go
// Below are equivalent;
example := 123
ptr := &example
// AND
ptr := ptrs.Int(123)
// AND
ptr := ptrs.Value(123)
```

``` go
// Below are equivalent;
var example bool
ptr := &example
// AND
ptr := ptrs.Bool(false)
// AND
ptr := ptrs.False()
```

``` go
// Below are equivalent;
type Priority uint8
const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)
example := PriorityHigh
ptr := &example
// AND
ptr := ptrs.Value(PriorityHigh)
```

``` go
// Below are equivalent;
var example Priority
ptr := &example
// AND
ptr := ptrs.Zero[Priority]()
```

There's many more functions available to explore!

## Issues

If you have any problems or would like to see changes currently in development you can do so
[here](https://github.com/neocotic/go-pointers/issues).

## Contributors

If you want to contribute, you're a legend! Information on how you can do so can be found in
[CONTRIBUTING.md](https://github.com/neocotic/go-pointers/blob/main/CONTRIBUTING.md). We want your suggestions and pull
requests!

A list of contributors can be found in [AUTHORS.md](https://github.com/neocotic/go-pointers/blob/main/AUTHORS.md).

## License

Copyright © 2023 neocotic

See [LICENSE.md](https://github.com/neocotic/go-pointers/raw/main/LICENSE.md) for more information on our MIT license.
