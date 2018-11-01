/*
 * USAGE: ./codewars <password>
 *
 * curl localhost:8080/<password>         // CORRECT
 * curl localhost:8080/<not-the-password> // WRONG
 *
 */

package main

import (
  "os"
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path[1:] == os.Args[1] {
    fmt.Fprintf(w, "CORRECT\n")
  } else {
    fmt.Fprintf(w, "WRONG\n")
  }
}

func main() {
  fmt.Println("starting server...")
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
