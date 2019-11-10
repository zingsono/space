package main

import (
	"log"

	"github.com/toolkits/net"
)

func main() {
	s, _ := net.IntranetIP()
	log.Println(s)
}
