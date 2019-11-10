package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"golang.org/x/net/websocket"

	"config/graph"
	"config/mgodb"
)

func main() {
	fmt.Println("******************************************************************************************")
	defer PrintStack()
	args()
	app.ListenAndServe()
}

func args() {
	var (
		// 服务名与端口号
		name string
		port int

		// Mongodb连接字符串
		mgdb string
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.StringVar(&mgdb, "mgdb", "", "Mongodb Connection String")
	flag.Parse()
	if mgdb == "" {
		log.Fatalf("Error: Params '-mgdb' can not be empty > %s", os.Args)
	}
	app.Name = name
	app.Port = port
	// -mgdb=mongodb://msde:msde0508@121.40.83.200:37017/msde?authSource=msde&authMechanism=SCRAM-SHA-1
	mgodb.SetDatabase("", mgdb)
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))

	if err := recover(); err != nil {
		log.Fatalf("** Main Fatalf-> %s", err)
	}
}

var app = &Application{Name: "config", Host: "0.0.0.0", Port: 5800}

type Application struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (app *Application) ListenAndServe() {
	app.Handles()
	var defaultHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(404) })

	log.Printf("MicroService: %s   ListenAndServe %s:%d", app.Name, app.Host, app.Port)
	log.Printf("Start server http://127.0.0.1:%d", app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), defaultHandler))
}

func (app *Application) Handles() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("MicroService:config"))
	})

	// Graphql服务
	http.Handle("/graphql", graph.GraphqlHttpHandler())

	// 配置订阅服务
	http.Handle("/ws/config", websocket.Handler(func(ws *websocket.Conn) {
		for {
			var mq map[string][]string
			err := websocket.JSON.Receive(ws, &mq)
			if err != nil {
				log.Panicln(err)
			}
			// 启动新线程，watch mongodb collection
			go func() {
				mgodb.NewConfig().Watch(nil, func(value map[string]interface{}) {
					err := websocket.JSON.Send(ws, value)
					if err != nil {
						log.Panicln(err)
					}
				})
			}()
		}
	}))

}
