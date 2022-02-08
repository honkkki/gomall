package main

import (
	"flag"
	"fmt"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/honkkki/gomall/code/mall/service/user/api/internal/config"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/handler"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case errorx.CodeError:
			return http.StatusOK, e.Response()
		default:
			exi := errorx.NewCodeError(errorx.InternalError, e.Error())
			ex, ok := exi.(errorx.CodeError)
			if !ok {
				return http.StatusInternalServerError, e.Error()
			}
			return http.StatusOK, ex.Response()
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
