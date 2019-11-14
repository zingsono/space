package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("EEG--------------------------------------------------------")
	log.Println("** Server start....")

	// 默认首页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("micro-control")
		io.WriteString(w, "version 1")
	})

	// TODO 设置登录会话安全验证
	// Graphql服务
	// http.Handle("/graphql", cors.Default().Handler(hgraph.GraphqlHttpHandler()))

	// HttpServer
	err := http.ListenAndServe(":50508", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("END--------------------------------------------------------")
}
