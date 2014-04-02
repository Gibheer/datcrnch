package filestore

import "os"
import "fmt"
import data "github.com/gibheer/datcrnch/data"
import encoding "encoding/binary"

func OpenForWrite(filename string) *os.File {
  f, _ := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0600)
  return f
}

func OpenForRead(filename string) *os.File {
  f, _ := os.Open(filename)
  return f
}

func Write(file *os.File, data *data.DataItem) {
  err := encoding.Write(file, encoding.BigEndian, int32(23))
  if err != nil {
    fmt.Println("Something went wrong:", err)
  }
}

func Read(file *os.File) int32 {
  var i int32
  err := encoding.Read(file, encoding.BigEndian, &i)
  if err != nil {
    fmt.Println("Something went wrong:", err)
  }
  return i
}
