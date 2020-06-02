package model

import (
	"github.com/SmartHobbyjd/gerenciador/shared/database"
)

type EditRes struct {
	Produtos     []Produtos
	Clientes     []Clientes
	Expenses     []Expenses
	Pedido       int
	Datapedido   string
	Produto      string
	Quantidade   int
	Valorunidade float64
	Valortotal   float64
	Datafabi     string
	Dataentrega  string
	Cliente      string
	Status       string
	Statuspedido string
	Obs          string
}

func Edit(nPedido string) EditRes {
	editres := EditRes{}
	// Get the URL `?pedido=X` parameter

	err := database.SQL.Get(&editres, "SELECT * FROM ordens WHERE pedido=?", nPedido)
	if err != nil {
		panic(err.Error())
	}

	// Create a slice to store all data from struct
	res := []Clientes{}
	productos := []Produtos{}
	err = database.SQL.Select(&res, "SELECT * FROM clientes")
	if err != nil {
		panic(err.Error())
	}

	err = database.SQL.Select(&productos, "SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	editres.Produtos = productos
	editres.Clientes = res
	return editres
}
