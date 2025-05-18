# expresso

A small library for parsing math expressions.

## Features

- Basic operations
- Handle parentheses
- Respect operator precedence
- Unary operators
- Whitespace tolerance
- Error handling
- Input validation

## Installation

```bash
go get github.com/J-Me-2307/expresso
```

## Usage

```go
package main

import (
 "fmt"

 "github.com/J-Me-2307/expresso/expresso"
)

func main() {
 expression := "4 * (19 + 6)"

  // In a real use case you should always handle the error gracefuly
 res, err := expresso.Evaluate(expression)
 if err != nil {
  panic(err)
 }

 fmt.Println(res) // Output: 100
}

```
