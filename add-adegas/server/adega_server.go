package server

import (
	"context"
	"errors"

	"github.com/Israel-Ferreira/add-adegas/models"
	"github.com/Israel-Ferreira/add-adegas/services/pb"
	"gorm.io/gorm"
)

type AdegaServiceServer struct {
	pb.UnimplementedAddAdegaServiceServer
	db *gorm.DB
}

func (ads *AdegaServiceServer) AddAdega(context context.Context, pbReq *pb.AdegaRequestBody) (*pb.AddAdegaResponse, error) {

	if pbReq.Nome == "" {
		return &pb.AddAdegaResponse{}, errors.New("Erro: O nome da adega deve estar preenchido")
	}

	if pbReq.Email == "" {
		return &pb.AddAdegaResponse{}, errors.New("Erro: o email da adega deve estar preenchido")
	}

	adega := models.NewAdega(pbReq)

	txn := ads.db.Save(&adega)

	if err := txn.Error; err != nil {
		return &pb.AddAdegaResponse{}, err
	}

	return &pb.AddAdegaResponse{Msg: "Adega cadastrada com sucesso", Status: 201}, nil
}

func NewAdegaServer(db *gorm.DB) *AdegaServiceServer {
	return &AdegaServiceServer{
		db: db,
	}
}
