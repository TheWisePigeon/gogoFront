package handler

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) (result string) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
  result = "Sap nigga"
  return
}
