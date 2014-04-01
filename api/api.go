package api

import (
  "net/http"
  "fmt"
  data "github.com/gibheer/datcrnch/data"
)

func Listen(hostname string, port int) {
  http.HandleFunc("/", handler)
  http.ListenAndServe(fmt.Sprintf("%s:%d", hostname, port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    fmt.Fprintf(w, "This is so wrong")
    return
  }
  d := data.CreateDataItem("foo", 23)
  fmt.Fprint(w, d)
}
