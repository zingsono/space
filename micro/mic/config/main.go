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
	"config/mgdb"
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
		mongo string
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.StringVar(&mongo, "mgdb", "", "Mongodb Connection String")
	flag.Parse()
	if mongo == "" {
		log.Fatalf("Error: Params '-mgdb' can not be empty > %s", os.Args)
	}
	app.Name = name
	app.Port = port
	// -mgdb=mongodb://msde:msde0508@121.40.83.200:37017/msde?authSource=msde&authMechanism=SCRAM-SHA-1
	mgdb.SetConnectString(mongo)
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
		// 保存单项配置对应的ws连接集合,map[k1]map[k2]*websocket.Conn, k1=订阅服务名 k2=当前连接服务名
		confws sync.Map
	)
	pushConf := func(kv map[string]interface{}) {
		for k, v := range kv {
			smap, ok := confws.Load(k)
			if !ok {
				return
			}
			for name, ws := range smap.(map[string]*websocket.Conn) {
				err := websocket.JSON.Send(ws, kv)
				if err != nil {
					log.Print(err)
				}
				log.Printf("给服务<%s>推送配置信息成功 %s=%s", name, k, v)
			}
		}

	}
	// Watch config
	go mgdb.NewConfig().Watch(pushConf)

	// 配置订阅服务
	http.Handle("/ws/config", websocket.Handler(func(ws *websocket.Conn) {
		log.Printf("request uri %s", ws.Request().RequestURI)
		// 订阅当前服务的服务名
		name := ws.Request().FormValue("name")

		defer ws.Close()
		for {
			var arr []string
			err := websocket.JSON.Receive(ws, &arr)
			if err != nil {
				log.Printf("ws连接断开 %s", err)
				break
			}
			log.Printf("接收消息：%s", arr)

			// 当前ws连接，使用map缓存
			for _, v := range arr {
				smap, ok := confws.Load(v)
				var m map[string]*websocket.Conn
				if ok {
					m = smap.(map[string]*websocket.Conn)
				} else {
					m = make(map[string]*websocket.Conn)
				}
				m[name] = ws
				confws.Store(v, m)
			}
			// 查询数据库回复订阅配置信息
			mgdb.NewConfig().Query(arr, pushConf)
		}
	}))

}
