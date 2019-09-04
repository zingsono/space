package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rs/cors"

	"mic/config"
	"mic/model"
)

func main() {
	log.Println("EEG--------------------------------------------------------")
	conf := config.NewConfig()
	log.Printf("** Server start http://127.0.0.1:%d", conf.Server.Port)

	// 默认首页
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "version 1")
	})

	// TODO 设置登录会话安全验证
	// Graphql服务
	http.Handle("/graphql", cors.Default().Handler(model.GraphqlHttpHandler))

	// HttpServer
	err := http.ListenAndServe(fmt.Sprintf(":%d", conf.Server.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("END--------------------------------------------------------")
}
