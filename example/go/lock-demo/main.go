// lock应用demo
package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func main() {
	fmt.Println("start lock main")
	mutex.Lock()
	fmt.Println("get locked main")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println("start lock ", i)
			mutex.Lock()
			defer mutex.Unlock()
			fmt.Println("get locked ", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock main")
	mutex.Unlock()
	fmt.Println("get unlocked main")
	time.Sleep(time.Second)
}
