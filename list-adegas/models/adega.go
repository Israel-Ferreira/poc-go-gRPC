package models

import "gorm.io/gorm"

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
