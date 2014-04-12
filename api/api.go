package api

import (
  "time"
  "fmt"
  "net/http"
  "strconv"
  encoding "encoding/binary"
  filestore "github.com/gibheer/datcrnch/filestore"
)

type Api struct {
  filename string
}

func (a *Api) Handler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    r.ParseForm()

    var d filestore.RawDataPoint
    i, err := strconv.ParseInt(r.PostForm["value"][0], 0, encoding.Size(d.Value))

    if err != nil {
      fmt.Fprint(w, "Something has gone very wrong:", err)
      return
    }
    d.Value = int32(i)

    d.Key = "foo"
    d.Timestamp = time.Now()
    a.WriteData(&d)
    fmt.Println("write successful")
  } else {
    var d filestore.RawDataPoint

    a.ReadData(&d)

    fmt.Fprint(w, "Found Data!", d)
  }
}

func (a *Api) WriteData(d filestore.DataPoint) {
  f := filestore.OpenForWrite(a.filename)
  defer f.Close()

  d.Write(f)
}

func (a *Api) ReadData(d filestore.DataPoint) {
  f := filestore.OpenForRead(a.filename)
  defer f.Close()

  stat, _  := f.Stat()
  filesize := stat.Size()
  fmt.Println(filesize, d.Size(), filesize / d.Size())

  err := d.Read(f)
  if err != nil {
    fmt.Println("Error", err)
  }
}

func Listen(hostname string, port int, filename string) {
  api := Api{filename}

  http.HandleFunc("/server/", api.Handler)
  bind := fmt.Sprintf("%s:%d", hostname, port)
  fmt.Println("Starting Server on", bind)
  http.ListenAndServe(bind, nil)
}
