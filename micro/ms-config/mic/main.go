package main

import (
	"mic/hh/server"
	"mic/model"
)

func main() {
	for i := 0; i < 500; i++ {
		// time.Sleep(1*time.Second)
		go func() {
			(&model.MsConfig{}).FindOne("config")
		}()
	}
	server.Run()
}
