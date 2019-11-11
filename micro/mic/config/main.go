package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"sync"

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

type Application struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

var app = &Application{Name: "config", Host: "0.0.0.0", Port: 5800}

func (app *Application) ListenAndServe() {
	app.Handles()
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

	var (
		// 服务被哪些服务订阅
		confarr sync.Map
		// WS连接缓存
		confws sync.Map
	)
	pushConf := func(key string, value map[string]interface{}) {
		arrLoad, ok := confarr.Load(key)
		if !ok {
			return
		}
		for _, item := range arrLoad.([]string) {
			wsLoad, ok := confws.Load(item)
			if !ok {
				continue
			}
			err := websocket.JSON.Send(wsLoad.(*websocket.Conn), value)
			if err != nil {
				log.Print(err)
			}
		}
	}
	// Watch config
	go mgodb.NewConfig().Watch(pushConf)

	// 配置订阅服务
	http.Handle("/ws/config", websocket.Handler(func(ws *websocket.Conn) {
		log.Printf("request uri %s", ws.Request().RequestURI)
		for {
			var mq map[string][]string
			err := websocket.JSON.Receive(ws, &mq)
			if err != nil {
				log.Panicln(err)
			}
			log.Printf("接收消息：%s", mq)
			for k, v := range mq {
				for _, n := range v {
					load, ok := confarr.Load(n)
					if ok {
						confarr.Store(n, append(load.([]string), k))
					} else {
						confarr.Store(n, []string{k})
					}
				}
				confws.Store(k, ws)

				// 查询数据库回复订阅配置信息
				mgodb.NewConfig().Query(v, pushConf)
			}
		}
	}))

}
