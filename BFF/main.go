package main

import (
	"log"
	"net"

	"github.com/atoyr/MagiaRecord/BFF/app"
	"github.com/atoyr/MagiaRecord/BFF/app/services"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("client connection error:")
	}
	server := grpc.NewServer()
	mgService := &services.MagicalGirlService{}
	app.RegisterGetMagicalGirlServiceServer(server, mgService)
	server.Serve(listenPort)
}
