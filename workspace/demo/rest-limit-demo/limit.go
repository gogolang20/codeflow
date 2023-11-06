package main

import (
	"flag"
	"fmt"

	"codeflow/workspace/demo/rest-limit-demo/internal/config"
	"codeflow/workspace/demo/rest-limit-demo/internal/handler"
	"codeflow/workspace/demo/rest-limit-demo/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/limit.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// hey -z 1s -c 90 -q 1 'http://localhost:8888/ping'
// hey -z 1s -c 110 -q 1 'http://127.0.0.1:8888/ping'
