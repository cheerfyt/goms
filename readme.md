golang ms package
=================

** copy from node ms package **

- install: `go get -v -u github.com/tate-fan/goms`
- test: `make test`

- usage:

  ```go
  package main

  import (
      "fmt"
      "github.com/tate-fan/goms"
  )

  func main() {
      value, err := goms.Ms("100d")

      if err != nil {
          panic(err.Error())
      }

      fmt.Printf("%.2f", value)
  }
  ```