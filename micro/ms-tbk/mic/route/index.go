package route

import (
	"net/http"
)

func init() {

	// 默认首页
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "version 1")
	})

	// 默认首页
	http.HandleFunc("/index2", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "version 2")
	})

}
