# go-log

Simple Golang logger package with multi-platform color support, with 3 different levels of logs (info, warn, error)

# How to use
```go
package main

import "github.com/desperatee/go-log"

func main() {
  logger := log.NewLogger()
  logger.Log(log.Info, "Hello world!")
}
