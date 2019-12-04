package main

import (
	"gateway/graph"
)

func main() {
	graph.ListenServe()

	/*http.ListenAndServe(":8989",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,"ok")
	}))*/
}
