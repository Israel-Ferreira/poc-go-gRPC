package main

import (
	"log"
	"net"

	"github.com/Israel-Ferreira/add-adegas/config"
	"github.com/Israel-Ferreira/add-adegas/models"
	"github.com/Israel-Ferreira/add-adegas/server"
	"github.com/Israel-Ferreira/add-adegas/services/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9000")

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

	db.AutoMigrate(&models.Adega{}, &models.Endereco{})

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAddAdegaServiceServer(grpcServer, server.NewAdegaServer(db))

	log.Println("Servidor gRPC iniciado na porta 9000")
	grpcServer.Serve(l)

}
