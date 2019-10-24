// Http Server
package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"mic/conf"
	"mic/graph"
)

func Post(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(405)
			w.Write([]byte("仅支持POST请求"))
		}
	})
}

func PostMapping(path string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	http.Handle(path, Post(http.HandlerFunc(handlerFunc)))
}

func Run() {
	server := conf.Now.Server
	log.Printf("** Server start http://127.0.0.1:%d", server.Port)
	// Graphql Server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.Port), cors.Default().Handler(graph.GraphqlHttpHandler)))
}
