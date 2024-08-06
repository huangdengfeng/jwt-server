package test

import (
	log "github.com/sirupsen/logrus"
	"jwt-server/entity/config"
	gclient "jwt-server/entity/grpc/client"
	"jwt-server/entity/grpc/server"
	"jwt-server/entity/pb"
	"jwt-server/service"
	"os"
	"testing"
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
	var createServer = func() pb.JwtServer {
		return &service.JwtServerImpl{}
	}
	server.Start(createServer())

	// start client
	client = gclient.CreateClient(config.Global.Server.Listen)

	log.Infof("[test] set up init success")
}

func teardown() {
	config.Shutdown()
	log.Infof("[test] tear down success")
}
