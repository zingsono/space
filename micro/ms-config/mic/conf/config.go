// 当前服务配置文件内容
package conf

type Config struct {
	Application *Application `json:"application"`
	Server      *Server      `json:"server"`
	Mongo       *Mongo       `json:"mongo"`
}

type Application struct {
	Name string `json:"name"`
}

type Server struct {
	Port int `json:"port"`
}

// 对应Mongo服务
type Mongo struct {
	Db0 string `json:"db0"`
}

// 当前配置内容
var Now = &Config{
	Application: &Application{Name: "config"},
	Server:      &Server{Port: 10508},
	Mongo:       &Mongo{Db0: "mongodb://test:test@121.40.83.200:37017/test?authSource=admin&authMechanism=SCRAM-SHA-1"},
}

// 读取当前配置
/*func Now() *Config {
	return NowFile("")
}

func NowFile(filename string) *Config {
	if filename == "" {
		filename = "./config.json"
	}
	if nowConfig != nil {
		return nowConfig
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	nowConfig := &Config{}
	err = json.Unmarshal(data, nowConfig)
	if err != nil {
		log.Fatalf("解析配置文件'%s'出错 %s", filename, err.Error())
	}
	return nowConfig
}*/
