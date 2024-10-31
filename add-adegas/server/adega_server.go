package server

import (
	"context"
	"errors"

	"github.com/Israel-Ferreira/add-adegas/services/pb"
)

type AdegaServiceServer struct {
	pb.UnimplementedAddAdegaServiceServer
}

func (ads *AdegaServiceServer) AddAdega(context context.Context, pbReq *pb.AdegaRequestBody) (*pb.AddAdegaResponse, error) {

	if pbReq.Nome == "" {
		return &pb.AddAdegaResponse{}, errors.New("Erro: O nome da adega deve estar preenchido")
	}

	if pbReq.Email == "" {
		return &pb.AddAdegaResponse{}, errors.New("Erro: o email da adega deve estar preenchido")
	}

	return &pb.AddAdegaResponse{}, nil
}

func NewAdegaServer() *AdegaServiceServer {
	return &AdegaServiceServer{}
}
