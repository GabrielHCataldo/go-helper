Go Helper
=================
<!--suppress ALL -->
<img align="right" src="gopher-helper.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.0.0-vividgreen.svg)](https://github.com/GabrielHCataldo/go-helper/releases/tag/v1.0.5)
[![Go Report Card](https://goreportcard.com/badge/github.com/GabrielHCataldo/go-helper)](https://goreportcard.com/report/github.com/GabrielHCataldo/go-helper)
[![Coverage Status](https://coveralls.io/repos/GabrielHCataldo/go-helper/badge.svg?branch=main&service=github)](https://coveralls.io/github/GabrielHCataldo/go-helper?branch=main)
[![Open Source Helpers](https://www.codetriage.com/gabrielhcataldo/go-helper/badges/users.svg)](https://www.codetriage.com/gabrielhcataldo/go-helper)
[![GoDoc](https://godoc.org/github/GabrielHCataldo/go-helper?status.svg)](https://pkg.go.dev/github.com/GabrielHCataldo/go-helper/helper)
![License](https://img.shields.io/dub/l/vibe-d.svg)

[//]: # ([![build workflow]&#40;https://github.com/GabrielHCataldo/go-helper/actions/workflows/go.yml/badge.svg&#41;]&#40;https://github.com/GabrielHCataldo/go-helper/actions&#41;)

[//]: # ([![Source graph]&#40;https://sourcegraph.com/github.com/go-helper/helper/-/badge.svg&#41;]&#40;https://sourcegraph.com/github.com/go-helper/helper?badge&#41;)

[//]: # ([![TODOs]&#40;https://badgen.net/https/api.tickgit.com/badgen/github.com/GabrielHCataldo/go-helper/helper&#41;]&#40;https://www.tickgit.com/browse?repo=github.com/GabrielHCataldo/go-helper&#41;)

The go-helper project came to facilitate the use of validations, conversions, formatting in your go project with
incredible flexibility and simple to use. Below we list some implemented features:

- No more using != or == nil, use better described functions for that.
- No more using len to check if it is empty or not.
- Use robust validation functions ready to be used.
- Check all possible types of values in a simple and intuitive way.
- Check single or multiple values if they are empty, null, the same, etc. without having to write lines and lines of
  code.
- Use new structure validations easily.
- Help us to be the best and most complete golang utility library and help the community.

Installation
------------

Use go get.

	go get github.com/GabrielHCataldo/go-helper

Then import the go-helper package into your own code.

```go
import "github.com/GabrielHCataldo/go-helper/helper"
```

Usability and documentation
------------
**IMPORTANT**: Always check the documentation in the structures and functions fields.
For more details on the examples, visit [All examples link](https://github/GabrielHCataldo/go-helper/blob/main/_example)

### Empty and Nil validations

For empty checks we can insert any value of any type, see the examples:

- #### STRUCT

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
    "time"
)

type exampleStruct struct {
    Name      string
    BirthDate time.Time
    Map       map[string]any
}

func main() {
    nStruct := exampleStruct{}
    isEmpty := helper.IsEmpty(nStruct)
    logger.Info("struct is empty?", isEmpty)
    nStruct.Name = "test name"
    isNotEmpty := helper.IsNotEmpty(nStruct)
    logger.Info("struct is not empty?", isNotEmpty)
}
```

Outputs:

    [INFO 2024/01/03 20:03:06] main.go:18: struct is empty? true
    [INFO 2024/01/03 20:03:06] main.go:21: struct is not empty? true

- #### SLICE

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    nSlice := []any{}
    isEmpty := helper.IsEmpty(nSlice)
    logger.Info("slice is empty?", isEmpty)
    nSlice = append(nSlice, "test")
    isNotEmpty := helper.IsNotEmpty(nSlice)
    logger.Info("slice is not empty?", isNotEmpty)
}
```

Outputs:

    [INFO 2024/01/03 20:52:20] main.go:12: slice is empty? true
    [INFO 2024/01/03 20:52:20] main.go:15: slice is not empty? true


- #### MAP AND POINTER MAP

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    nMap := map[string]any{}
    isEmpty := helper.IsEmpty(&nMap)
    logger.Info("pointer map is empty?", isEmpty)
    nMap["test"] = "value string"
    isNotEmpty := helper.IsNotEmpty(nMap)
    logger.Info("map is not empty?", isNotEmpty)
}
```

Outputs:

    [INFO 2024/01/03 20:58:44] main.go:12: pointer map is empty? true
    [INFO 2024/01/03 20:58:44] main.go:15: pointer map is not empty? true

- #### STRING

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    nBlankString := "  "
    isEmpty := helper.IsEmpty(nBlankString)
    logger.Info("string is empty?", isEmpty)
    nBlankString = "test name"
    isNotEmpty := helper.IsNotEmpty(nBlankString)
    logger.Info("string is not empty?", isNotEmpty)
}
```

Outputs:

    [INFO 2024/01/03 20:48:24] main.go:12: string is empty? true
    [INFO 2024/01/03 20:48:24] main.go:15: string is not empty? true

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go).

### Convert

### Format

### Validate

### For more examples

- [Empty](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go)
- [Equal](https://github/GabrielHCataldo/go-helper/blob/main/_example/equal/main.go)
- [Float](https://github/GabrielHCataldo/go-helper/blob/main/_example/float/main.go)
- [Format](https://github/GabrielHCataldo/go-helper/blob/main/_example/format/main.go)
- [Int](https://github/GabrielHCataldo/go-helper/blob/main/_example/int/main.go)
- [String](https://github/GabrielHCataldo/go-helper/blob/main/_example/string/main.go)
- [Time](https://github/GabrielHCataldo/go-helper/blob/main/_example/time/main.go)
- [Type](https://github/GabrielHCataldo/go-helper/blob/main/_example/type/main.go)
- [Validate](https://github/GabrielHCataldo/go-helper/blob/main/_example/validate/main.go)
- [Validator](https://github/GabrielHCataldo/go-helper/blob/main/_example/validator/main.go)

Thanks
------
- https://github.com/go-playground/validator/v10
- https://github.com/klassmann/cpfcnpj
- https://github.com/nyaruka/phonenumbers
- https://github.com/leekchan/accounting

How to contribute
------
Make a pull request, or if you find a bug, open it
an Issues.

License
-------
Distributed under MIT license, see the license file within the code for more details.