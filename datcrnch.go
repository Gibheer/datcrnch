package main

import (
  "fmt"
  filestore "github.com/gibheer/datcrnch/filestore"
//  api "github.com/gibheer/datcrnch/api"
  data "github.com/gibheer/datcrnch/data"
)

func main() {
//  api.Listen("127.0.0.1", 9876)
  f := filestore.Open("foo")
  data := data.CreateDataItem("foo", 23)
  filestore.Write(f, data)
  f.Close()
  fmt.Println(f)
}
