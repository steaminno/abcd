package model

import (
	"fmt"

	"github.com/SmartHobbyjd/gerenciador/shared/database"
)

type ClientProduct struct {
	Produtos []Produtos
	Clientes []Clientes
	Expenses []Expenses
}

func New() ClientProduct {
	// db := dbConn()
	// n := Clientes{}

	// Create a slice to store all data from struct
	clients := []Clientes{}
	//producto := Produtos{}
	productos := []Produtos{}
	//selDB, err := database.Query("SELECT * FROM clientes")
	err := database.SQL.Select(&clients, "SELECT id, pedido, cellnumber, fname, lname, address, status, obs FROM clientes")
	if err != nil {
		panic(err.Error())
	}

	err = database.SQL.Select(&productos, "SELECT id, valorunidade, product_cost, pname, pdescription, pdatafabi, pimg, status, obs FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	result := ClientProduct{}
	result.Clientes = clients
	result.Produtos = productos
	//fmt.Printf("MIME Header: %+v\n", result)
	return result
}

func Insert(ordens *Ordens) {

	// Open database connection

	//_, err := database.SQL.Exec("INSERT INTO note (content, user_id) VALUES (?,?)", content, userID)
	_, err := database.SQL.Exec("INSERT INTO ordens(datapedido, produto, quantidade, valorunidade, valortotal, datafabi, dataentrega, cliente, status, statuspedido, obs) VALUES(?,?,?,?,?,?,?,?,?,?,?)", ordens.Datapedido, ordens.Produto, ordens.Quantidade, ordens.Valorunidade, ordens.Valortotal, ordens.Datafabi, ordens.Dataentrega, ordens.Cliente, ordens.Status, ordens.Statuspedido, ordens.Obs)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	// Execute the prepared SQL, getting the form fields
	//insForm.Exec(ordens.Datapedido, ordens.Produto, ordens.Quantidade, ordens.Valorunidade, ordens.Valortotal, ordens.Datafabi, ordens.Dataentrega, ordens.Cliente, ordens.Status, ordens.Statuspedido, ordens.Obs)
	// Redirect to HOME
}
