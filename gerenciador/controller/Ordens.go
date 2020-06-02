package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func Index(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)
	result := model.Index()
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "index/Index"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Orden"] = result.Orden
	v.Vars["Valortotal"] = FormatCurrencyFloat(result.Valortotal, 2)
	v.Vars["TotalOrden"] = result.TotalOrden
	v.Vars["TotalClient"] = result.TotalClient
	v.Vars["TotalProduct"] = result.TotalProduct
	v.Vars["TotalExpense"] = FormatCurrencyInt(result.TotalExpense)
	v.Vars["TotalSaled"] = FormatCurrencyFloat(result.TotalSaled, 2)
	v.Vars["TotalSaledpaid"] = result.TotalSaledpaid
	v.Vars["TotalSaledopen"] = result.TotalSaledopen
	v.Vars["TotalSaledparcial"] = result.TotalSaledparcial
	v.Vars["TotalSellsPartial"] = FormatCurrencyFloat(result.TotalSellsPartial, 2)
	v.Vars["TotalSellsOpen"] = FormatCurrencyFloat(result.TotalSellsOpen, 0)

	fmt.Println(v)

	v.Render(w)

}

//TotalSellsOpen:R$ %!f(int=100) TotalSellsPartial:R$ %!f(int=100) Valortotal:R$ %!f(int=100)

func FormatCurrencyFloat(valor float64, dec int) string {
	p := message.NewPrinter(language.BrazilianPortuguese)

	formatZ := "R$ %.0f"

	format2 := "R$ %.2f"

	if dec == 0 {
		return p.Sprintf(formatZ, valor)
	}
	return p.Sprintf(format2, valor)
}

func FormatCurrencyInt(valor int) string {
	p := message.NewPrinter(language.BrazilianPortuguese)
	return p.Sprintf("R$ %.2f", float64(valor/100))
}

func Show(w http.ResponseWriter, r *http.Request) {
	nPedido := r.URL.Query().Get("pedido")
	sess := session.Instance(r)
	result := model.Show(nPedido)
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "show/Show"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Orden"] = result

	v.Render(w)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Insert 1")

	ordens := model.Ordens{}

	ordens.Datapedido = r.FormValue("datapedido")
	ordens.Produto = r.FormValue("produto")
	ordens.Quantidade, _ = strconv.Atoi(r.FormValue("quantidade"))
	ordens.Valorunidade, _ = strconv.ParseFloat(r.FormValue("valorunidade"), 64)
	ordens.Valortotal, _ = strconv.ParseFloat(r.FormValue("valortotal"), 64)
	ordens.Datafabi = r.FormValue("datafabi")
	ordens.Dataentrega = r.FormValue("dataentrega")
	ordens.Cliente = r.FormValue("cliente")
	ordens.Status = r.FormValue("status")
	ordens.Statuspedido = r.FormValue("statuspedido")
	ordens.Obs = r.FormValue("obs")
	fmt.Println(" Insert 2")
	model.Insert(&ordens)
	fmt.Println(" Insert 3")
	// Redirect to HOME
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Update 1")
	if r.Method == "POST" {
		fmt.Println(" Update 2")
		ordens := model.Ordens{}

		ordens.Datapedido = r.FormValue("datapedido")
		ordens.Produto = r.FormValue("produto")
		ordens.Quantidade, _ = strconv.Atoi(r.FormValue("quantidade"))
		ordens.Valorunidade, _ = strconv.ParseFloat(r.FormValue("valorunidade"), 64)
		ordens.Valortotal, _ = strconv.ParseFloat(r.FormValue("valortotal"), 64)
		ordens.Datafabi = r.FormValue("datafabi")
		ordens.Dataentrega = r.FormValue("dataentrega")
		ordens.Cliente = r.FormValue("cliente")
		ordens.Status = r.FormValue("status")
		ordens.Statuspedido = r.FormValue("statuspedido")
		ordens.Obs = r.FormValue("obs")
		ordens.Pedido, _ = strconv.Atoi(r.FormValue("upedido"))
		fmt.Println(" Update 3")
		model.Update(ordens)
		fmt.Println(" Update 4")

	}
	fmt.Println(" Update 5")
	// Redirect to Home
	http.Redirect(w, r, "/", 301)
}

// Function Delete destroys a row based on Pedido
func Delete(w http.ResponseWriter, r *http.Request) {

	nPedido := r.URL.Query().Get("pedido")

	model.Delete(nPedido)

	// Redirect a HOME
	http.Redirect(w, r, "/", 301)
}
