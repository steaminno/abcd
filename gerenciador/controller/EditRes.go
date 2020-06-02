package controller

import (
	"net/http"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func Edit(w http.ResponseWriter, r *http.Request) {

	nPedido := r.URL.Query().Get("pedido")
	sess := session.Instance(r)
	result := model.Edit(nPedido)
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "edit/Edit"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["editres"] = result

	v.Render(w)

}
