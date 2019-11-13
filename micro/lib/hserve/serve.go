package hserve

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// 监听服务，可在handlerFun中加载业务配置
func ListenAndServe(handlerFun func(app *Application)) {
	fmt.Println("******************************************************************************************")
	defer printStack()
	args()
	handlerFun(app)
	app.listenAndServe()
}

func args() {
	var (
		// 服务名与端口号
		name string
		port int
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))
	flag.Parse()
	app.Name = name
	app.Port = port
}

func printStack() {
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

var app = &Application{Name: "default", Host: "0.0.0.0", Port: 55800}

func (app *Application) listenAndServe() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}
		w.Write([]byte("MicroService:config"))
	})

	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}
