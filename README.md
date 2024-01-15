Go Helper
=================
<!--suppress ALL -->
<img align="right" src="gopher-helper.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.1.0-vividgreen.svg)](https://github.com/GabrielHCataldo/go-helper/releases/tag/v1.1.0)
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
- You don't need to use more lens to check if it is empty or not.
- Use robust validation functions out of the box.
- Check all types of possible values in a simple and intuitive way.
- Check single or multiple values if they are empty, null, equal, etc. without having to write lines and lines of
  code.
- Easy and intuitive rounding.
- Format and Convert for you to use in the best way in your projects.
- Use new framework validations easily.
- Help us be the best and most complete golang utility library and help the community.

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

### Empty validate

For empty checks we can insert any value of any type, see the examples:

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
    isNotEmpty := helper.IsNotEmpty(&nStruct)
    logger.Info("pointer struct is not empty?", isNotEmpty)
}
```

Outputs:

    [INFO 2024/01/03 20:03:06] main.go:18: struct is empty? true
    [INFO 2024/01/03 20:03:06] main.go:21: struct is not empty? true

You can validate multiple values simultaneously, whether they are empty or not, see:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    var anyPointer *any
    var nMap map[string]any
    var nSlice []any
    var nString string
    var nInt int
    var nFloat float64
    var nBool bool
   
    allEmpty := helper.AllEmpty(anyPointer, nMap, nSlice, nString, nInt, nFloat, nBool)
    logger.Info("all are empty?", allEmpty)
   
    anyPointer = helper.ConvertToPointer(any("value string"))
    nMap = map[string]any{
      "test": "value",
    }
    nString = "value"
    nInt = 1
    nFloat = 1.0
    nBool = true
    nSlice = append(nSlice, any("value"))
   
    allNotEmpty := helper.AllNotEmpty(anyPointer, nMap, nString, nInt, nFloat, nBool)
    logger.Info("all are not empty?", allNotEmpty)
}
```

Outputs:

    [INFO 2024/01/04 05:09:56] main.go:17: all are empty? true
    [INFO 2024/01/04 05:09:56] main.go:27: all are not empty? true

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go).

### Nil validate
Validating the "nil" value is simple, but be careful, as
only the types **Pointer**, **Map**, **Chan**, **Interface**, **Slice** and **Array**
can have a nil value in go, see an example below:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    var anyPointer *any
    isNil := helper.IsNil(anyPointer)
    logger.Info("pointer is nil?", isNil)
    anyPointer = helper.ConvertToPointer(any("value string"))
    isNotNil := helper.IsNotNil(anyPointer)
    logger.Info("pointer is not nil?", isNotNil)
}
```

Outputs:

    [INFO 2024/01/04 05:01:57] main.go:12: pointer is nil? true
    [INFO 2024/01/04 05:01:57] main.go:15: pointer is not nil? true

You can also be validating multiple values simultaneously, see the example below:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    var anyPointer *any
    var nMap map[string]any
    var nSlice []any
    var nChan chan struct{}
    var nInterface interface{}
    
    allNil := helper.AllNil(anyPointer, nMap, nSlice, nChan, nInterface)
    logger.Info("all are nil?", allNil)
    
    anyPointer = helper.ConvertToPointer(any("value string"))
    nMap = map[string]any{}
    nSlice = append(nSlice, any("value"))
    nChan = make(chan struct{}, 1)
    nInterface = "value"
    
    allNotNil := helper.AllNotNil(anyPointer, nMap, nSlice, nChan, nInterface)
    logger.Info("all are not nil?", allNotNil)
}
```

Outputs

    [INFO 2024/01/04 05:19:48] main.go:17: all are nil? true
    [INFO 2024/01/04 05:19:48] main.go:26: all are not nil? true

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go).

### IsEqual
Checking equals is very simple, and works for any type and number 
of parameters entered, if it is **Pointer**, go-helper will obtain the
real value for comparison, see one of the examples:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
)

func main() {
    s1 := "value"
    s2 := "value"
    s3 := "value"
    s4 := "value"
    
    equals := helper.IsEqual(s1, s2, s3, s4)
    logger.Info("equals?", equals)
    
    s1 = "value1"
    
    notEquals := helper.IsNotEqual(s1, s2, s3, s4)
    logger.Info("not equals?", notEquals)
}
```

Outputs:

    [INFO 2024/01/04 05:32:42] main.go:15: equals? true
    [INFO 2024/01/04 05:32:42] main.go:20: not equals? true

### Type
With go-helper it is very simple to know the type of your variable, see the example:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "github.com/GabrielHCataldo/go-logger/logger"
    "time"
)

func main() {
    var a any = map[string]any{}
    
    // for negative use IsNotMap
    isMap := helper.IsMap(a)
    logger.Info("is map?", isMap)
    
    a = "value1"

    // for negative use IsNotString
    isString := helper.IsString(a)
    logger.Info("is string?", isString)
    
    a = 12
	
    // for negative use IsNotInt
    isInt := helper.IsInt(a)
    logger.Info("is int?", isInt)
    
    a = true

    // for negative use IsNotBool
    isBool := helper.IsBool(a)
    logger.Info("is bool?", isBool)
    
    a = time.Now()

    // for negative use IsNotTime
    isTime := helper.IsTime(a)
    logger.Info("is time?", isTime)
}
```

Outputs:

    [INFO 2024/01/04 05:47:38] main.go:13: is map? true
    [INFO 2024/01/04 05:47:38] main.go:18: is string? true
    [INFO 2024/01/04 05:47:38] main.go:23: is int? true
    [INFO 2024/01/04 05:47:38] main.go:28: is bool? true
    [INFO 2024/01/04 05:47:38] main.go:33: is time? true

### For more examples

- [Empty](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go)
- [IsEqual](https://github/GabrielHCataldo/go-helper/blob/main/_example/equal/main.go)
- [Float](https://github/GabrielHCataldo/go-helper/blob/main/_example/float/main.go)
- [Format](https://github/GabrielHCataldo/go-helper/blob/main/_example/format/main.go)
- [Convert](https://github/GabrielHCataldo/go-helper/blob/main/_example/convert/main.go)
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