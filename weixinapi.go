package main

import (
	"flag"
	"fmt"
	"github.com/Holyson/test-go-zero-cors/rest/httpx"
	"net/http"

	"cors-test/internal/config"
	"cors-test/internal/handler"
	"cors-test/internal/svc"

	"github.com/Holyson/test-go-zero-cors/core/conf"
	"github.com/Holyson/test-go-zero-cors/rest"
)

var configFile = flag.String("f", "etc/weixinapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//conf.MustLoad(*configFile, &c)
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.Error(w, err)
	}), rest.WithCors("*"))
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
