package router

import (
	"log"
	"net/http"
)

func RunHttpServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Gateway Index"))
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:50801", nil))
}
