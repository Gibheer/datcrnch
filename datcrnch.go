package main

import (
//  "fmt"
  filestore "github.com/gibheer/datcrnch/filestore"
  api "github.com/gibheer/datcrnch/api"
  "time"
)

func main() {
  f := filestore.OpenForWrite("foo")
  defer f.Close()
  go api.Listen("127.0.0.1", 9876, f)

  time.Sleep(5 * time.Minute)
}
