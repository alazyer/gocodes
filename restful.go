package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
    ws.Path("/v1")
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)

	ws2 := new(restful.WebService)
    ws2.Path("/v1")
	ws2.Route(ws2.GET("/hello").To(hello2))
	restful.Add(ws2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}

func hello2(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world2")
}
