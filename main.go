package main

import (
    "log"
    "net/http"
)

func main() {
  port := 3000
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening on port:", port)
  http.ListenAndServe(":3000", nil)
}
