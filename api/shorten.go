package handler

import (
	"fmt"
	"net/http"
)

func Shorten(w http.ResponseWriter, r *http.Request){
	var longUrl string = ""
	fmt.Println(longUrl)
	fmt.Fprintf(w, "<h1>SIUUUUUUUUUUUUU</h1>")
}