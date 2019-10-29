// Http Server
package route

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rs/cors"

	"config/conf"
	"config/graph"
)

/*func Post(h http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(405)
			w.Write([]byte("仅支持POST请求"))
		}
	})
}*/

func Run() {
	server := conf.Now.Server
	log.Printf("** Server start http://127.0.0.1:%d", server.Port)

	// 默认首页
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "version 1")
	})

	// Graphql Server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", server.Port), cors.Default().Handler(graph.GraphqlHttpHandler())))
}
