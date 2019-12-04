package graph

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func ListenServe() {
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
	)
	flag.StringVar(&name, "name", app.Name, fmt.Sprintf("Set Application name. Default '%s'", app.Name))
	flag.IntVar(&port, "port", app.Port, fmt.Sprintf("Set Port. Default is %d", app.Port))

	flag.Parse()

	app.Name = name
	app.Port = port
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
	Name    string `json:"name"`
	Version string `json:"version"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

var app = &Application{Name: "gateway", Version: "0.0.1", Host: "", Port: 7908}

func (app *Application) ListenAndServe() {
	Handles()
	log.Printf("MicroService: %s  ListenAndServe %s:%d   Start server http://127.0.0.1:%d", app.Name, app.Host, app.Port, app.Port)
	panic(http.ListenAndServe(fmt.Sprintf("%s:%d", app.Host, app.Port), nil))
}

func Handles() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*if r.URL.Path != "/" {
			http.NotFoundHandler().ServeHTTP(w, r)
			return
		}*/
		io.WriteString(w, "MicroService:"+app.Name)
	})

	http.HandleFunc("/tt", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "TT MicroService:"+app.Name)
	})

	h := GraphqlHttpHandler()
	// Graphql服务
	http.Handle("/graphql", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := context.WithValue(r.Context(), "token", "============================r.Header.Get()")
		// token、用户id、用户名等信息存储header，用于记录操作日志
		// h.ContextHandler(ctx, w, r)

		/*bytes,err := ioutil.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w,err.Error())
			return
		}
		body := string(bytes)
		log.Print("Body="+body)*/
		// io.WriteString(w,body)

		h.ServeHTTP(w, r)
		// io.WriteString(w,"gggg")
	}))

	http.HandleFunc("/api/v2/gateway", func(w http.ResponseWriter, r *http.Request) {
		/*bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		source := source.NewSource(&source.Source{
			Body: bytes,
			Name: "GraphQL request",
		})
		// parse the source
		AST, err := parser.Parse(parser.ParseParams{Source: source})
		if err != nil {
			log.Print(err)
		}
		// 解析出一级字段作为服务名
		for _, operationDefinition := range AST.Definitions {
			selectionSet := operationDefinition.(*ast.OperationDefinition).SelectionSet

			for _, selection := range selectionSet.Selections {
				field := selection.(*ast.Field).Name.Value

				log.Print("sss=" + field)
			}
		}*/

		io.WriteString(w, "OK")
	})
	http.HandleFunc("/api/v2/gateway2", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		hql := string(bytes)
		// 遍历查询字符串,拆分不同服务
		//services := make(map[string]string)

		// 计数器
		lb := 0

		for i, ch := range hql {
			char := fmt.Sprintf("%c", ch)
			log.Printf("%d = %s ", i, char)
			if char == "{" {
				lb++
			}
			if char == "}" {
				lb--
				if lb == 1 {

				}
			}

		}

		io.WriteString(w, "OK")
	})

}
