package main

import (
	log "github.com/sirupsen/logrus"
	"jwt-server/entity/config"
	"jwt-server/entity/grpc/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()
	defer config.Shutdown()

	s := server.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	o := <-sig
	log.Printf("recieve signal %s ,server will stop gracefully", o.String())
	server.Stop(s)
}
