/*package main

import (
    "io"
    "log"
    "net/http"
)

func HelloGoServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, this is a GoServer")
}

func main() {
    http.HandleFunc("/", HelloGoServer)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServer ", err)
    }
}*/


package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()

  rh := http.RedirectHandler("http://www.baidu.com", 307)
  mux.Handle("/foo", rh)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}