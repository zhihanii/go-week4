package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kserver "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	video "go-tiktok/app/kitex_gen/video/videoservice"
	"go-tiktok/app/pkg/bound"
	"go-tiktok/app/pkg/conf"
	"go-tiktok/app/pkg/constants"
	"go-tiktok/app/pkg/middleware"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

// 第四周作业
// 这是我最近在做的项目, 使用Kitex框架搭建rpc服务

func main() {
	c := conf.Default()

	// 使用etcd注册
	r, err := etcd.NewEtcdRegistry([]string{c.Etcd.Addr})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	if err != nil {
		panic(err)
	}

	s := video.NewServer(NewVideoService(c),
		kserver.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}),
		kserver.WithMiddleware(middleware.CommonMiddleware),
		kserver.WithMiddleware(middleware.ServerMiddleware),
		kserver.WithServiceAddr(addr),
		kserver.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		kserver.WithMuxTransport(),
		kserver.WithSuite(trace.NewDefaultServerSuite()),
		kserver.WithBoundHandler(bound.NewCpuLimitHandler()),
		kserver.WithRegistry(r), // 使用服务注册中间件
	) // 在启动Server时会进行服务注册

	srv := &server{s}

	var srvs = []Server{srv}
	var sigs = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	ctx, cancel := context.WithCancel(context.Background())
	var opts = []Option{
		WithCancel(ctx, cancel),
		WithTimeout(time.Second),
		WithSrvs(srvs),
		WithSigs(sigs),
	}

	app := new(App)
	for _, opt := range opts {
		opt(app) // 配置App
	}

	// 启动
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
