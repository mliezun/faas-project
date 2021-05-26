package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello Go!")
}

func main() {
	r := router.New()
	r.GET("/", Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8899"
	}

	fmt.Println("Listening on port", port)

	log.Fatal(fasthttp.ListenAndServe(":"+port, r.Handler))
}
