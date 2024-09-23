package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpc"
	"net/http"
	"time"

	"exporter/internal/config"
	"exporter/internal/handler"
	"exporter/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/monitor-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	go RefreshCache(ctx)
	go Report(ctx)
	server.Start()
}
func RefreshCache(ctx *svc.ServiceContext) {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	for true {
		fmt.Println("开始刷新")
		time.Sleep(time.Second)
		urlPath := fmt.Sprintf("http://localhost:%d/refresh/refresh", ctx.Config.Port)
		resp, _ := httpc.Do(context.Background(), http.MethodGet, urlPath, nil)
		if resp == nil {
			continue
		}
		fmt.Println("结束刷新", resp)
		fmt.Println(resp.Body.Close())
		time.Sleep(time.Second * 50)
	}
}

func Report(ctx *svc.ServiceContext) {
	defer func() {
		if e := recover(); e != nil {
			return
		}
	}()
	for true {
		fmt.Println("开始报告")
		time.Sleep(time.Second * 4)
		urlPath := fmt.Sprintf("http://localhost:%d/monitor/report", ctx.Config.Port)
		resp, _ := httpc.Do(context.Background(), http.MethodPost, urlPath, nil)
		if resp == nil {
			continue
		}
		fmt.Println("结束报告", resp)
		fmt.Println(resp.Body.Close())
		time.Sleep(time.Minute * 10)
	}
}
