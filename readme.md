golang ms package
=================

**node ms package golang complete**

- install: `go get -v -u github.com/cheerfyt/goms`
- test: `go test ./`

- usage:

  ```go
  package main

  import (
      "fmt"
      "github.com/cheerfyt/goms"
  )

  func main() {
      value, err := goms.Parse("100d")

      if err != nil {
          panic(err.Error())
      }

      fmt.Printf("%.2f", value)
  }
  ```