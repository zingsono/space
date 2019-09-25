package main

import (
	"context"
	"fmt"
	"time"
)

func A1(ctx context.Context) {
	fmt.Println("func A1 ", ctx.Value("a"))
}

func B1(ctx context.Context) {
	fmt.Println("func B1 ", ctx.Value("b"))
}

func UB1(ctx context.Context) {
	context.WithValue(ctx, "b", "b-value-u")
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "a", "a-value")
	ctx = context.WithValue(ctx, "b", "b-value")

	time.Sleep(2 * time.Second)
	go A1(ctx)
	go UB1(ctx)
	time.Sleep(2 * time.Second)
	go B1(ctx)

	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}
