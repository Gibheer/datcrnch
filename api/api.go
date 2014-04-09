package api

import (
  "time"
  "net/http"
  "fmt"
  "os"
  "strconv"
  filestore "github.com/gibheer/datcrnch/filestore"
  encoding "encoding/binary"
)

type Api struct {
  file *os.File
}

func (a *Api) Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "This is working!")
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
  a.WriteData(d)
  fmt.Println("write successful")
}

func (a *Api) WriteData(d filestore.DataPoint) {
  d.Write(a.file)
}

func Listen(hostname string, port int, file *os.File) {
  api := Api{file}

  http.HandleFunc("/server/", api.Handler)
  bind := fmt.Sprintf("%s:%d", hostname, port)
  fmt.Println("Starting Server on", bind)
  http.ListenAndServe(bind, nil)
}
