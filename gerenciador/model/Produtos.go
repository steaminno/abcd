package model

import (
	"fmt"

	"github.com/SmartHobbyjd/gerenciador/shared/database"
)

type Produtos struct {
	ID           int
	Pname        string
	Pdescription string
	Valorunidade float64
	Product_cost float64
	Pdatafabi    string
	Pimg         string
	Status       string
	Obs          string
}

func Product() []Produtos {

	res := []Produtos{}

	err := database.SQL.Select(&res, "SELECT * FROM produtos ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	return res

}

func ProductStore(pr Produtos) {
	_, err := database.SQL.Exec("INSERT INTO produtos(pname, pdescription, valorunidade, product_cost, pdatafabi, pimg, status, obs) VALUES(?,?,?,?,?,?,?,?)", pr.Pname, pr.Pdescription, pr.Valorunidade, pr.Product_cost, pr.Pdatafabi, pr.Pimg, pr.Status, pr.Obs)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
}
