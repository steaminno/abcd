package model

import (
	"fmt"
	"time"

	"github.com/SmartHobbyjd/gerenciador/shared/database"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Ordens struct {
	Pedido       int
	Datapedido   string
	Produto      string
	Quantidade   int
	Valorunidade float64
	//Product_cost int
	Valortotal   float64
	Datafabi     string
	Dataentrega  string
	Cliente      string
	Status       string
	Statuspedido string
	Obs          string
}

func (o Ordens) FormatCurrencyFloat(valor float64) string {
	p := message.NewPrinter(language.BrazilianPortuguese)
	return p.Sprintf("R$ %.2f", valor)
}

func (o Ordens) FormatCurrencyInt(valor int) string {
	p := message.NewPrinter(language.BrazilianPortuguese)
	return p.Sprintf("R$ %.2f", float64(valor/100))
}

func (o Ordens) FormatDate(d string) string {
	fmt.Println("date is ", d)
	nd, e := time.Parse("2006-01-02T15:04:05Z", d)

	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("date is ", d)
	fmt.Println("format date is ", nd.Format("02/01/2006"))

	return nd.Format("02/01/2006")
}

func Index() Res {

	// db := dbConn()
	// n := Clientes{}

	// Create a slice to store all data from struct
	res := []Ordens{}

	//selDB, err := database.Query("SELECT * FROM clientes")
	err := database.SQL.Select(&res, "SELECT  * FROM ordens ORDER BY pedido DESC")
	if err != nil {
		panic(err.Error())
	}

	reses := Res{}
	reses.Orden = res
	reses.Valortotal, reses.TotalOrden, reses.TotalClient, reses.TotalProduct, reses.TotalExpense, reses.TotalSaledpaid, reses.TotalSaledopen, reses.TotalSaledparcial, reses.TotalSaled, reses.TotalSellsPartial, reses.TotalSellsOpen = getamount()
	return reses
}

func getamount() (Valortotal float64, TotalOrden int, TotalClient int, TotalProduct int, TotalExpense int, TotalSaledpaid int, TotalSaledopen int, TotalSaledparcial int, TotalSaled float64, TotalSellsPartial float64, TotalSellsOpen float64) {

	err := database.SQL.Get(&Valortotal, "SELECT SUM(valortotal) FROM ordens WHERE status ='pago'")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalClient, "SELECT COUNT(id) FROM clientes")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalExpense, "SELECT COUNT(id) FROM expenses")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalOrden, "SELECT COUNT(pedido) FROM ordens")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalProduct, "SELECT COUNT(id) FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSaledpaid, "SELECT COUNT(pedido) FROM ordens WHERE status = 'pago'")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSaledopen, "SELECT COUNT(pedido) FROM ordens WHERE status = 'em aberto'")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSaledparcial, "SELECT COUNT(pedido) FROM ordens WHERE status = 'parcial'")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSaled, "SELECT SUM(valortotal) FROM ordens")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSellsPartial, "SELECT SUM(valortotal) FROM ordens WHERE status = 'parcial'")
	if err != nil {
		panic(err.Error())
	}
	err = database.SQL.Get(&TotalSellsOpen, "SELECT SUM(valortotal) FROM ordens WHERE status = 'em aberto'")
	if err != nil {
		panic(err.Error())
	}
	//SELECT COUNT(pedido) FROM ordens
	return
}

/*

	res := Ordens{}

	err := database.SQL.Select(&res, "SELECT * FROM ordens WHERE pedido=?", nPedido)
	if err != nil {
		panic(err.Error())
	}

	reses := Res{}
	reses.Orden = res
	return reses


*/

func Show(nPedido string) Ordens {
	orden := Ordens{}

	err := database.SQL.Get(&orden, "SELECT * FROM ordens WHERE pedido=?", nPedido)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}

	return orden

}

func Update(o Ordens) {

	_, err := database.SQL.Exec("UPDATE ordens SET datapedido=?, produto=?, quantidade=?, valorunidade=?, valortotal=?, datafabi=?, dataentrega=?, cliente=?, status=?, statuspedido=?, obs=? WHERE pedido=?",
		o.Datapedido, o.Produto, o.Quantidade, o.Valorunidade, o.Valortotal, o.Datafabi, o.Dataentrega, o.Cliente, o.Status, o.Statuspedido, o.Pedido)
	if err != nil {
		panic(err.Error())
	}
}

// Function Delete destroys a row based on Pedido
func Delete(id string) {

	_, err := database.SQL.Exec("DELETE FROM ordens WHERE pedido=?", id)
	if err != nil {
		panic(err.Error())
	}
}
