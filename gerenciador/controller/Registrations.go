package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func Registration(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)
	res := model.Registration()
	v := view.New(r)
	v.Name = "registration/Registration"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Registration"] = res

	v.Render(w)

}

func RegistrationCreate(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)

	v := view.New(r)
	v.Name = "registration/RegistrationCreate"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)

	v.Render(w)

}

func RegistrationStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in RegistrationStore 1")
	// Open database connection
	reg := model.Registrations{}

	// Check the request form METHOD
	if r.Method == "POST" {
		fmt.Println("in RegistrationStore 2")
		// Get the values from Form
		reg.Username = r.FormValue("username")
		reg.Password = r.FormValue("password")
		reg.Email = r.FormValue("email")
		reg.Cellnumber, _ = strconv.Atoi(r.FormValue("cellnumber"))
		reg.Status = r.FormValue("status")
		reg.Role = r.FormValue("role")
		reg.Obs = r.FormValue("obs")
		fmt.Println("in RegistrationStore 3")
		model.RegistrationStore(reg)
		fmt.Println("in RegistrationStore 4")
	}
	fmt.Println("in RegistrationStore 5")
	// Redirect to Registration
	//http.Redirect(w, r, "/registration", 301)
	// sess := session.Instance(r)

	// v := view.New(r)
	// v.Name = "registration/Registration"
	// v.Vars["token"] = csrfbanana.Token(w, r, sess)

	// v.Render(w)
	http.Redirect(w, r, "/registration", http.StatusPermanentRedirect)
}
