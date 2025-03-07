package server

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"jwt-server/entity/config"
	"jwt-server/entity/pb"
	"net"
	"time"
)

var kp = keepalive.ServerParameters{
	MaxConnectionIdle: 3 * time.Minute, // 如果连接空闲超过这个时间，服务端将关闭连接
}
var kep = keepalive.EnforcementPolicy{
	PermitWithoutStream: true, // 空闲时候发ping ，而不是断开连接
}

func Start(jwtServer pb.JwtServer) *grpc.Server {
	var grpcServer = grpc.NewServer(grpc.KeepaliveParams(kp), grpc.KeepaliveEnforcementPolicy(kep), createDefaultInterceptor())
	pb.RegisterJwtServer(grpcServer, jwtServer)

	go func() {
		listen, err := net.Listen("tcp", config.Global.Server.Listen)
		if err != nil {
			log.Fatalf("listen error [%s]", err)
		}
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("server serve error [%s]", err)
		}
	}()
	return grpcServer
}

func Stop(grpcServer *grpc.Server) {
	grpcServer.GracefulStop()
}
