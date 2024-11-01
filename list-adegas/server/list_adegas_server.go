package server

import (
	"context"
	"log"

	"github.com/Israel-Ferreira/list-adegas/models"
	"github.com/Israel-Ferreira/list-adegas/services/pb"
	"gorm.io/gorm"
)

type ListAdegasServer struct {
	pb.UnimplementedListAndGetAdegasServiceServer

	db *gorm.DB
}

func (las *ListAdegasServer) ListAdegas(pbReq *pb.ListAdegasFilter, streamGrpc pb.ListAndGetAdegasService_ListAdegasServer) error {

	var adegas []models.Adega

	txn := las.db.Preload("EnderecoAdega").Find(&adegas)

	if err := txn.Error; err != nil {
		log.Println("Erro ao fazer a consulta no banco de dados: ", err)
		return err
	}

	log.Println()

	for _, adega := range adegas {
		adegaResp := pb.AdegaResponse{
			Id:    int32(adega.ID),
			Nome:  adega.Nome,
			Email: adega.Email,
			Endereco: &pb.EnderecoResponse{
				Id:          int32(adega.EnderecoAdega.ID),
				Logradouro:  adega.EnderecoAdega.Logradouro,
				Numero:      adega.EnderecoAdega.Numero,
				Bairro:      adega.EnderecoAdega.Bairro,
				Complemento: adega.EnderecoAdega.Complemento,
				Cidade:      adega.EnderecoAdega.Cidade,
				Uf:          adega.EnderecoAdega.Uf,
				Cep:         adega.EnderecoAdega.Cep,
			},
		}

		if err := streamGrpc.Send(&adegaResp); err != nil {
			log.Println("Erro: ", err)
			return err
		}
	}

	return nil
}

func (las *ListAdegasServer) GetAdegaById(ctx context.Context, pbReq *pb.GetAdegasByIdFilter) (*pb.AdegaResponse, error) {
	return &pb.AdegaResponse{}, nil
}

func NewListAdegaServer(db *gorm.DB) *ListAdegasServer {
	return &ListAdegasServer{
		db: db,
	}
}
