# Grettings in Go

This package lets a very simple form to get Grettings in Go.

## Installation
Execute this command for install this package:

```bash
go get github.com/ferbj/greetings
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/ferbj/greetings"
)

func main() {
    message, err := greetings.Hello("Fernando")
    if err != nil {
        fmt.Println("There was an error:", err)
        return
    }
}