package serve

import (
	"log"
	"strings"
	"sync"

	"golang.org/x/net/websocket"
)

var (
	DiscoverName string

	// 注册数据通过内存存储 map[serviceName]map[id]Serve
	instanceMap sync.Map
	// 服务被谁订阅 map[watchServiceName]map[id]*websocket.Conn
	watch sync.Map

	REG MsgType = 1
	OFF MsgType = 0
)

// 1=注册  0=剔除
type MsgType int

// 服务注册结构体
type Instance struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`

	Watch []string `json:"watch"`

	msgType MsgType `json:"type"`

	// 当前ws连接
	Ws *websocket.Conn `json:"-"`

	// 标记是否是同步消息
	Sync bool
}

func (s *Instance) Type(reg func(), off func()) {
	if s.msgType == REG {
		reg()
	}
	if s.msgType == OFF {
		off()
	}
}

// 服务信息推送
func (s *Instance) push() {
	discoverServiceName, ws := DiscoverName, s.Ws
	// 保存注册数据
	lm, ok := instanceMap.Load(s.Name)
	if !ok {
		lm = make(map[string]*Instance)
	}
	s.Type(func() { lm.(map[string]*Instance)[s.Id] = s }, func() { delete(lm.(map[string]*Instance), s.Id) })

	// 保存ws连接
	for _, v := range s.Watch {
		m, ok := watch.Load(v)
		if !ok {
			m = make(map[string]*websocket.Conn)
		}
		s.Type(func() { m.(map[string]*websocket.Conn)[s.Id] = ws }, func() { delete(m.(map[string]*websocket.Conn), s.Id) })
	}

	// 给当前服务发送消息
	s.Type(func() {
		for _, v := range s.Watch {
			lms, ok := instanceMap.Load(v)
			if ok {
				for _, s := range lms.(map[string]*Instance) {
					err := websocket.JSON.Send(ws, s)
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	}, func() {})

	// 给所有订阅通道发送消息
	wsm, ok := watch.Load(s.Name)
	if ok {
		for _, v := range wsm.(map[string]*websocket.Conn) {
			err := websocket.JSON.Send(v, s)
			if err != nil {
				log.Print(err)
			}
		}
	}

	s.syncInstance(discoverServiceName)
}

// 发送消息
func (s *Instance) Send() {
	s.msgType = REG
	s.push()
}

// 服务关闭
func (s *Instance) Close() {
	s.msgType = OFF
	s.push()
}

// 同步实例数据
func (s *Instance) syncInstance(name string) {
	if s.Sync {
		return
	}
	s.Sync = true
	wsn, ok := watch.Load(name)
	if ok {
		for _, v := range wsn.(map[string]*websocket.Conn) {
			err := websocket.JSON.Send(v, s)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

func Panic(err error) {
	log.Print(err)
}

// 服务注册
func Reg(dc string, instance *Instance) {
	for _, url := range strings.Split(dc, ",") {
		if url == "" {
			continue
		}
		go func() {
			ws, err := websocket.Dial(url, "", "http://127.0.0.1:10000")
			if err != nil {
				log.Print(err)
				return
			}
			err = websocket.JSON.Send(ws, instance)
			if err != nil {
				log.Print(err)
				return
			}
			for {
				var msg *Instance
				websocket.JSON.Receive(ws, &msg)
				if err != nil {
					log.Print(err)
					break
				}
				// 成功收到消息，客户端缓存
				value, ok := instanceMap.Load(msg.Name)
				if !ok {
					value = make(map[string]*Instance)
				}
				value.(map[string]*Instance)[msg.Id] = msg
				instanceMap.Store(msg.Name, value)
			}
		}()
	}
}
