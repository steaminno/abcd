package model

import (
	"github.com/SmartHobbyjd/gerenciador/shared/database"
)

type Clientes struct {
	ID         int
	Pedido     int
	Fname      string
	Lname      string
	Cellnumber int
	Address    string
	Status     string
	Obs        string
}

func Client() []Clientes {

	res := []Clientes{}
	err := database.SQL.Select(&res, "SELECT * FROM clientes ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	return res

}

func ClientStore(cl Clientes) {

	database.SQL.Exec("INSERT INTO clientes(pedido, fname, lname, cellnumber, address, status, obs) VALUES(?,?,?,?,?,?,?)", cl.Pedido, cl.Fname, cl.Lname, cl.Cellnumber, cl.Address, cl.Status, cl.Obs)

}
