package models

import (
	"github.com/Israel-Ferreira/add-adegas/services/pb"
	"gorm.io/gorm"
)

type Endereco struct {
	gorm.Model
	Logradouro  string
	Numero      string
	Complemento string
	Bairro      string
	Cidade      string
	Uf          string
	Cep         string
	AdegaID     uint
}

func NewEndereco(enderecoPbReq *pb.EnderecoRequest) Endereco {
	return Endereco{
		Logradouro:  enderecoPbReq.Logradouro,
		Bairro:      enderecoPbReq.Bairro,
		Complemento: enderecoPbReq.Complemento,
		Numero:      enderecoPbReq.Numero,
		Cep:         enderecoPbReq.Cep,
		Uf:          enderecoPbReq.Cep,
		Cidade:      enderecoPbReq.Cidade,
	}
}
