package models

import (
	"github.com/Israel-Ferreira/add-adegas/services/pb"
	"gorm.io/gorm"
)

type Adega struct {
	gorm.Model
	Nome     string
	Email    string
	Telefone string
	Ativa    bool

	HorarioFuncionamento string
	AceitaPedidosOnline  bool

	EnderecoAdega Endereco
}

func (adg *Adega) BeforeCreate(tx *gorm.DB) (err error) {
	adg.Ativa = true
	return
}

func NewAdega(adegaPbReq *pb.AdegaRequestBody) *Adega {
	return &Adega{
		Nome:                 adegaPbReq.Nome,
		Email:                adegaPbReq.Email,
		Telefone:             adegaPbReq.Telefone,
		AceitaPedidosOnline:  adegaPbReq.AceitaPedidosOnline,
		HorarioFuncionamento: adegaPbReq.HorarioFuncionamento,
		EnderecoAdega:        NewEndereco(adegaPbReq.Endereco),
	}
}
