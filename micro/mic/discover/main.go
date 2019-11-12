package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"golang.org/x/net/websocket"

	"discover/graph"
	"discover/serve"
)

func main() {
	fmt.Println("******************************************************************************************")
	defer PrintStack()
	args()
	app.ListenAndServe()
}

func args() {
	var (
		name string
		port int
		dc   string
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.StringVar(&dc, "dc", "", "Discover服务地址,多个使用逗号分隔，实例：ws://127.0.0.1:5801/discover,ws://127.0.0.1:5802/discover")
	flag.Parse()
	app.Name = name
	app.Port = port
	app.dc = dc
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

type Application struct {
	Name string
	Host string
	Port int

	dc string
}

var app = &Application{Name: "discover", Host: "0.0.0.0", Port: 5800}

func (app *Application) ListenAndServe() {
	app.Handles()

	// 注册服务
	serve.Reg(app.dc, &serve.Instance{
		Id:    app.Host + string(app.Port),
		Name:  app.Name,
		Host:  app.Host,
		Port:  app.Port,
		Watch: []string{app.Name},
	})

	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}

func (app *Application) Handles() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:config"))
	})

	// Graphql服务
	http.Handle("/graphql", graph.GraphqlHttpHandler())

	// 配置订阅服务
	http.Handle("/discover", websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		serve.DiscoverName = app.Name
		for {
			var instance *serve.Instance
			err := websocket.JSON.Receive(ws, &instance)
			instance.Ws = ws
			if err != nil {
				instance.Close()
				log.Printf("ws连接断开 %s", err)
				break
			}
			log.Printf("接收消息：%s", instance)
			instance.Send()
		}
	}))

}
