package controller

import (
	"net/http"
	"strconv"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func New(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)
	result := model.New()
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "client/New"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Produtos"] = result.Produtos
	v.Vars["Clientes"] = result.Clientes
	v.Render(w)

}

func Client(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	result := model.Client()
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "client/Client"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Clientes"] = result
	v.Render(w)

}

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)
	v := view.New(r)
	v.Name = "client/ClientCreate"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)

	v.Render(w)
	//tmpl.ExecuteTemplate(w, "ClientCreate", nil)
}

func ClientStore(w http.ResponseWriter, r *http.Request) {

	clientes := model.Clientes{}

	clientes.Pedido, _ = strconv.Atoi(r.FormValue("pedido"))
	clientes.Fname = r.FormValue("fname")
	clientes.Lname = r.FormValue("lname")
	clientes.Cellnumber, _ = strconv.Atoi(r.FormValue("cellnumber"))
	clientes.Address = r.FormValue("address")
	clientes.Status = r.FormValue("status")
	clientes.Obs = r.FormValue("obs")

	model.ClientStore(clientes)

	// Redirect to Client
	http.Redirect(w, r, "/client", 307)
}
