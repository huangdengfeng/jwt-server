package test

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"jwt-server/entity/config"
	c "jwt-server/entity/grpc/client"
	"jwt-server/entity/grpc/server"
	"jwt-server/entity/pb"
	"jwt-server/service"
	"os"
	"testing"
	"time"
)

var client pb.JwtClient

func TestMain(m *testing.M) {
	setup()
	// 运行测试
	exitCode := m.Run()
	// 退出测试
	teardown()
	os.Exit(exitCode)
}
func setup() {
	config.ServerConfigPath = "../conf"
	config.Init()
	var createJwtServer = func() pb.JwtServer {
		return &service.JwtServerImpl{}
	}
	server.Start(createJwtServer())

	// start client
	kacp := keepalive.ClientParameters{
		Time:                1 * time.Minute,  // 客户端每隔1min发送一次心跳ping
		Timeout:             10 * time.Second, // 如果没有收到服务端的心跳响应，认为连接失败的超时时间
		PermitWithoutStream: true,             // 即使没有活动的RPC流，也允许发送心跳
	}
	opts := []grpc.DialOption{grpc.WithKeepaliveParams(kacp), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: 1 * time.Second})}
	opts = append(opts, c.CreateDefaultInterceptor())
	conn, err := grpc.NewClient(config.Global.Server.Listen, opts...)

	if err != nil {
		log.Fatalf("connect error [%s]", err)
	}
	client = pb.NewJwtClient(conn)

	log.Infof("[test] set up init success")
}

func teardown() {
	config.Shutdown()
	log.Infof("[test] tear down success")
}
