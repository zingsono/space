package router

import (
	"net/http"
)

func init() {

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Home ..."))
	})

}
