package main

import (
	"flag"
	"fmt"
	"net/http"

	"server/internal/config"
	"server/internal/handler"
	"server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "net/http/pprof"
)

var configFile = flag.String("f", "etc/zero.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	RegisterProf()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func RegisterProf() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			panic("pprof server start error: " + err.Error())
		}
	}()
}
