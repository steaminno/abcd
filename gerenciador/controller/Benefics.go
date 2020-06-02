package controller

import (
	"net/http"
	"strconv"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func Benefic(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)
	res := model.Benefic()

	v := view.New(r)
	v.Name = "benefic/Benefic"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Benefic"] = res

	v.Render(w)

}

func BeneficCreate(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)
	res := model.BeneficCreate()
	v := view.New(r)
	v.Name = "benefic/BeneficCreate"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["BeneficCreate"] = res

	v.Render(w)

	//tmpl.ExecuteTemplate(w, "BeneficCreate", res)
}

func BeneficStore(w http.ResponseWriter, r *http.Request) {

	b := model.Benefics{}

	// Check the request form METHOD
	if r.Method == "POST" {

		// Get the values from Form
		b.Compname = r.FormValue("compname")
		b.Types = r.FormValue("types")
		b.Status = r.FormValue("status")
		b.Client_id, _ = strconv.Atoi(r.FormValue("cliente"))
		model.BeneficStore(b)
	}

	// Redirect to benefic
	http.Redirect(w, r, "/benefic", 307)
}
