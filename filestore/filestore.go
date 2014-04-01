package filestore

import "os"
import "fmt"
import data "github.com/gibheer/datcrnch/data"
import encoding "encoding/binary"

func Open(filename string) *os.File {
  f, _ := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0600)
  return f
}

func Write(file *os.File, data *data.DataItem) {
  fmt.Println("I'm writing, right?")
  err := encoding.Write(file, encoding.BigEndian, data.Value)
  if err != nil {
    fmt.Println("Something went wrong:", err)
  }
}
