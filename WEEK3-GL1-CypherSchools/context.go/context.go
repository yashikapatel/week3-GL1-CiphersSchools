package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go doSomething(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("timeline exceeded of 2 sec")
	}
	time.Sleep(time.Second * 3)
}
func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("Doing something bakwaas")

		}

	}
}

// func main() {
// 	ctx := context.Background()
// 	//seed some data in ctx
// 	//seedContext(ctx)
// 	ctx = seedContext(ctx)
// 	readCtx(ctx)

// }
func readCtx(ctx context.Context) {
	value := ctx.Value("one")
	fmt.Println(value)
}
func seedContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "one", "111")
	return ctx
}
