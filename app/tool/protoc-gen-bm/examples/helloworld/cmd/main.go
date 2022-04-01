package main

import (
	"flag"
	pb "go-common/app/tool/protoc-gen-bm/examples/helloworld/api"
	"go-common/app/tool/protoc-gen-bm/examples/helloworld/service"
	bm "go-common/library/net/http/blademaster"
	"log"
)

/*
*   curl http://127.0.0.1:8000/hello\?name\=world
*   curl 127.0.0.1:8000/echo  -H "Content-Type: application/json"  -X POST
 */
func main() {
	flag.Parse()
	engine := bm.NewServer(nil)
	// 注册 middleware 支持正则匹配
	engine.Inject("^/echo", func(c *bm.Context) {
		// do something
	})
	s := new(service.Service)
	pb.RegisterHelloBMServer(engine, s)
	if err := engine.Run("127.0.0.1:8000"); err != nil {
		log.Fatal(err)
	}
}
