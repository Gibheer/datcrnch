package main

import (
//  "fmt"
  api "github.com/gibheer/datcrnch/api"
  "time"
)

func main() {
  go api.Listen("127.0.0.1", 9876, "foo")

  time.Sleep(5 * time.Minute)
}
