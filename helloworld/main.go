package main

import (
	"fmt"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	_ "github.com/go-spring/starter-echo"
)

func main() {
	gs.GetMapping("/", func(ctx web.Context) {
		ctx.String("hello world!")
	})
	fmt.Println(gs.Run())
}
