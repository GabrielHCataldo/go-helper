Go Helper
=================
<!--suppress ALL -->
<img align="right" src="gopher-helper.png" alt="">

[![Project status](https://img.shields.io/badge/version-v1.7.0-vividgreen.svg)](https://github.com/GabrielHCataldo/go-helper/releases/tag/v1.7.0)
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

- No more using !=, ==, >, among others, start using better described functions for this.
- It is not necessary to use more len to check if it is empty or not, or greater than the same, just pass the values into intuitive functions.
- Use robust validation functions out of the box.
- Check all types of possible values in a simple and intuitive way.
- Check single or multiple values if they are empty, null, equal, etc. without having to write lines and lines of code.
- Easy rounding functions.
- Format and Convert functions for you to use in the best way in your projects, without having to write repetitive lines of code.
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
To check if the values are empty, just do as in the example below,
remembering that we also accept values of any type, see:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "log"
    "time"
)

type exampleStruct struct {
    Name      string
    BirthDate time.Time
    Map       map[string]any
}

func main() {
    // empty validations
    nStruct := exampleStruct{}
    nString := "  "
    isEmpty := helper.IsEmpty(nStruct, nString)
    log.Println("is empty?", isEmpty)
    nString = "test name"
    nStruct.Name = nString
    isNotEmpty := helper.IsNotEmpty(&nStruct, nString)
    log.Println("is not empty?", isNotEmpty)
	
    //return B value if A value is empty
    var a1 string
    var b1 string
    a1 = "a value is not empty"
    result := helper.IfEmptyReturns(a1, b1)
    log.Println("result not empty:", result)
	
    //return value not empty
    var a2 string
    var b2 string
    b2 = "b value is not empty"
    result = helper.ReturnNonEmptyValue(a2, b2)
    log.Println("result not empty:", result)
}
```

Outputs:

    is empty? true
    is not empty? true
    result not empty: a value is not empty
    result not empty: b value is not empty

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go).

### Nil validate
Validating “nil” values is simple, but be careful as
only the types **Pointer**, **Map**, **Chan**, **Interface** and **Slice**
can have a null value in go, see an example below:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "log"
)

func main() {
    var anyPointer *any
    var anotherValue []any
    isNil := helper.IsNil(anyPointer, anotherValue)
    log.Println("is nil?", isNil)
    anyPointer = helper.ConvertToPointer(any("value string"))
    anotherValue = []any{12, "test"}
    isNotNil := helper.IsNotNil(anyPointer, anotherValue)
    log.Println("is not nil?", isNotNil)
    
    //return B value if A value is nil
    var a1 *string
    var b1 string
    b1 = "b value is not nil"
    result1 := helper.IfNilReturns(a1, b1)
    log.Println("result not nil:", result1)
    
    //return value not nil
    var a2 *string
    var b2 *string
    v2 := "a value is not nil"
    a2 = &v2
    result2 := helper.ReturnNonNilValue(a2, b2)
    log.Println("result not nil:", result2)
}

```

Outputs:

    is nil? true
    is not nil? true
    result not nil: 0x14000223150
    result not nil: 0x1400025d7f0

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go).

### Equals validate
Checking equal is very simple, and works for any type and number 
of parameters entered, if it is **Pointer**, go-helper will obtain the
real value for comparison, see one of the examples:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "log"
)

func main() {
    s1 := "value"
    s2 := "value"
    s3 := "value"
    s4 := "value"
    
    equals := helper.Equals(s1, s2, s3, s4)
    log.Println("equals?", equals)
    
    s1 = "value1"
    
    notEquals := helper.IsNotEqualTo(s1, s2, s3, s4)
    log.Println("not equals?", notEquals)

    s1 = "VaLuE"
    s2 = "VaLuE"
    s3 = "VaLuE"
    s4 = "VaLuE"

    equals = helper.EqualsIgnoreCase(s1, s2, s3, s4)
    log.Println("equals ignore case?", equals)

    s1 = "value1"

    notEquals = helper.IsNotEqualToIgnoreCase(s1, s2, s3, s4)
    log.Println("not equals ignore case?", notEquals)
}
```

Outputs:

    equals? true
    not equals? true
    equals ignore case? true
    not equals ignore case? true

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/equal/main.go).

### Type validate
With go-helper it is very simple to know the type of your variable, see the example:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "log"
    "time"
)

func main() {
    var a any = map[string]any{}
    
    // for negative use IsNotMap
    isMap := helper.IsMapType(a)
    log.Println("is map?", isMap)
    
    a = "value1"

    // for negative use IsNotString
    isString := helper.IsStringType(a)
    log.Println("is string?", isString)
    
    a = 12
	
    // for negative use IsNotIntType
    isInt := helper.IsIntType(a)
    log.Println("is int?", isInt)
    
    a = true

    // for negative use IsNotBoolType
    isBool := helper.IsBoolType(a)
    log.Println("is bool?", isBool)
    
    a = time.Now()

    // for negative use IsNotTimeType
    isTime := helper.IsTimeType(a)
    log.Println("is time?", isTime)
}
```

Outputs:

    is map? true
    is string? true
    is int? true
    is bool? true
    is time? true

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/type/main.go).

### Convert
With go-helper you don't need to worry about laborious and repetitive conversions to implement, see a powerful 
example of conversion from any type A to any type B:

```go
package main

import (
    "github.com/GabrielHCataldo/go-helper/helper"
    "log"
    "time"
)

type exampleStruct struct {
    Name      string
    BirthDate time.Time
    Map       map[string]any
}

func main() {
    var destStruct exampleStruct
    // convert to struct to string base64, it's funny and sample
    s64 := helper.SimpleConvertToBase64(initExampleStruct())
    // convert string base64 to struct
    err := helper.ConvertToDest(s64, &destStruct)
    if helper.IsNotNil(err) {
        log.Println("error convert base64 to struct:", err)
        return
    }
    log.Println("converted base64 to struct:", destStruct)
    var destMap map[string]any
    // convert struct to simple map
    err = helper.ConvertToDest(destStruct, &destMap)
    if helper.IsNotNil(err) {
        log.Println("error convert struct to map:", err)
        return
    }
    log.Println("converted struct to map:", destMap)
}

func initExampleStruct() exampleStruct {
    return exampleStruct{
        Name:      "Foo Bar",
        BirthDate: time.Now(),
        Map:       map[string]any{"example": 1, "test": 2},
    }
}
```

Output:

    [INFO 2024/01/17 09:09:00] main.go:26: converted base64 to struct: {"Name":"Foo Bar","BirthDate":"2024-01-17T09:08:49-03:00","Map":{"example":1,"test":2}}
    [INFO 2024/01/17 09:09:00] main.go:34: converted struct to map: {"Name":"Foo Bar","BirthDate":"2024-01-17T09:08:49.038335-03:00","Map":{"example":1,"test":2}}

See other types of values as examples by accessing
the [link](https://github/GabrielHCataldo/go-helper/blob/main/_example/convert/main.go).

### For more examples

- [Empty](https://github/GabrielHCataldo/go-helper/blob/main/_example/empty/main.go)
- [Equal](https://github/GabrielHCataldo/go-helper/blob/main/_example/equal/main.go)
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