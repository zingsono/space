package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	hash32 := fnv.New32()
	hash32.Write([]byte("1234ajflkasjldkfj"))
	v := hash32.Sum32()
	fmt.Println(v % 32)

	fmt.Println(int32(32))

	var maps [10]int

	fmt.Println(maps[0])
	fmt.Println(maps[1])
}
