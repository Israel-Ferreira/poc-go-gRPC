package models

import (
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
