package main

import (
	"net/http"
)

func main() {

	s := "/Users/golang/golang"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s))))
	http.ListenAndServe(":8088", nil)
}
