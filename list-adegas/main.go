package main

import (
	"log"
	"net"

	"github.com/Israel-Ferreira/list-adegas/config"
	"github.com/Israel-Ferreira/list-adegas/server"
	"github.com/Israel-Ferreira/list-adegas/services/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9001")

	if err != nil {
		log.Fatalln(err)
	}

	godotenv.Load()

	dbEnvs := config.LoadDatabaseEnvs()

	db, err := config.DbOpenConnection(
		dbEnvs["host"],
		dbEnvs["dbUsername"],
		dbEnvs["dbPassword"],
		dbEnvs["dbName"],
		dbEnvs["dbPort"],
	)

	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterListAndGetAdegasServiceServer(grpcServer, server.NewListAdegaServer(db))

	log.Println("Servidor gRPC iniciado na porta 9001")
	grpcServer.Serve(l)

}
