package handler

import (
	"fmt"
	"net/http"
)

func Redirect(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>SIUUUUUUUUUUUUU</h1>")
}