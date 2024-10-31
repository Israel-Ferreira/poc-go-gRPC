package main

import (
	"log"
	"net"

	"github.com/Israel-Ferreira/add-adegas/server"
	"github.com/Israel-Ferreira/add-adegas/services/pb"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAddAdegaServiceServer(grpcServer, server.NewAdegaServer())

	grpcServer.Serve(l)

}
