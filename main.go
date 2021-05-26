package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/valyala/fasthttp"
)

type Proxy struct {
	client *fasthttp.HostClient
}

var proxyClients = []*Proxy{
	{
		client: &fasthttp.HostClient{
			Addr: "localhost:8080",
		},
	},
}

func ReverseProxyHandler(px *Proxy, ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	prepareRequest(req)
	if err := px.client.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	postprocessResponse(resp)
}

func prepareRequest(req *fasthttp.Request) {
	// do not proxy "Connection" header.
	req.Header.Del("Connection")
	// strip other unneeded headers.

	//fmt.Println(req.URI(), string(req.Header.RawHeaders()))

	// alter other request params before sending them to upstream host
}

func postprocessResponse(resp *fasthttp.Response) {
	// do not proxy "Connection" header
	resp.Header.Del("Connection")

	// strip other unneeded headers

	//fmt.Println(resp.RemoteAddr().String(), string(resp.Header.Header()))

	// alter other response data if needed
}

type opCode byte

const (
	opConstant opCode = iota
	opProxy
)

var constants = []interface{}{
	"Hello Go!",
}

var impls = []func(f *FunctionInstance){
	opConstant: func(f *FunctionInstance) {
		f.ip++
		f.reqCtx.WriteString(constants[f.proto.code[f.ip]].(string))
	},
	opProxy: func(f *FunctionInstance) {
		f.ip++
		ReverseProxyHandler(f.proto.proxies[f.proto.code[f.ip]], f.reqCtx)
	},
}

var code = []opCode{
	opConstant, 0,
}

type Function struct {
	code      []opCode
	proxies   []*Proxy
	constants []interface{}
}

func (f *Function) NewInstance(ctx *fasthttp.RequestCtx) *FunctionInstance {
	return &FunctionInstance{
		proto:  f,
		ip:     0,
		fp:     0,
		frame:  []interface{}{},
		reqCtx: ctx,
	}
}

type FunctionInstance struct {
	proto  *Function
	ip     int
	fp     int
	frame  []interface{}
	reqCtx *fasthttp.RequestCtx
}

func (f *FunctionInstance) Run() error {
	for f.ip = 0; f.ip < len(f.proto.code); f.ip++ {
		c := f.proto.code[f.ip]
		impls[c](f)
	}
	return nil
}

var testFn *Function = &Function{
	code:      code,
	proxies:   proxyClients,
	constants: constants,
}

type Hostname string
type RoutePattern string

type Route func(ctx *fasthttp.RequestCtx)

type Router map[RoutePattern]Route

var hostRepository = map[Hostname]Router{
	"forta.xyz": {
		"/": func(ctx *fasthttp.RequestCtx) {
			testFn.NewInstance(ctx).Run()
		},
	},
}

func Index(ctx *fasthttp.RequestCtx) {
	if host := ctx.URI().Host(); len(host) > 50 {
		ctx.SetStatusCode(http.StatusUnprocessableEntity)
		ctx.WriteString("Host too long")
	} else {
		router, ok := hostRepository[Hostname(host)]
		if !ok {
			ctx.SetStatusCode(http.StatusNotFound)
			ctx.WriteString("Not found")
		} else {
			for r, h := range router {
				if string(r) == string(ctx.Request.URI().Path()) {
					h(ctx)
					return
				}
			}
			ctx.SetStatusCode(http.StatusNotFound)
			ctx.WriteString("Not found")
		}
	}
}

func oldmain() {
	/*
		r := router.New()
		r.ANY("*", Index)
	*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8899"
	}

	fmt.Println("Listening on port", port)

	log.Fatal(fasthttp.ListenAndServe(":"+port, Index))
}
