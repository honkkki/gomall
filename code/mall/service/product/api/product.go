package main

import (
	"flag"
	"fmt"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/handler"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(errorx.ErrorHandler())

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
