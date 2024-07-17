package server

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jwt-server/entity/config"
	"jwt-server/entity/pb"
	"net"
)

func Start(jwtServer pb.JwtServer) *grpc.Server {
	var grpcServer = grpc.NewServer(CreateDefaultInterceptor())
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
