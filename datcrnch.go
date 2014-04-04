package main

import (
  "fmt"
  filestore "github.com/gibheer/datcrnch/filestore"
//  api "github.com/gibheer/datcrnch/api"
  "time"
)

func main() {
//  api.Listen("127.0.0.1", 9876)
  f := filestore.OpenForWrite("foo")

  values := filestore.IntValues{
    Count: 23,
    Average: 24,
    Min: 25,
    Max: 26,
    Percentile99:27,
  }
  d := filestore.AggregatedDataPoint{time.Now(), "foo", values}
  err := d.Write(f)
  if err != nil {
    fmt.Println("Bad:", err)
  }
  f.Close()

  f = filestore.OpenForRead("foo")
  var d2 filestore.AggregatedDataPoint
  err = d2.Read(f)
  if err != nil {
    fmt.Println("Bad 2:", err)
  }

  fmt.Println(d, d2)
}
