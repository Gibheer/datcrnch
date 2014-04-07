package api

import (
  "net/http"
  "fmt"
)

type Api struct {

}

func (a *Api) Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "This is working!")
}

func Listen(hostname string, port int) {
  api := Api{}

  http.HandleFunc("/server/", api.Handler)
  bind := fmt.Sprintf("%s:%d", hostname, port)
  fmt.Println("Starting Server on", bind)
  http.ListenAndServe(bind, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    fmt.Fprintf(w, "This is so wrong")
    return
  }
  r.ParseForm()
  fmt.Println(r.PostForm)
  fmt.Fprint(w, "This is working")
}
