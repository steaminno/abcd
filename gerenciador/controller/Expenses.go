package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/SmartHobbyjd/gerenciador/model"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/josephspurrier/csrfbanana"
)

func Expense(w http.ResponseWriter, r *http.Request) {

	expensres := model.Expense()

	fmt.Println("expensres ... ", expensres)
	fmt.Println("expensres.Expense ... ", expensres.Expense)

	sess := session.Instance(r)

	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "expense/Expense"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["expensres"] = expensres

	v.Render(w)

}

func ExpenseCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	res := model.ExpenseCreate()
	v := view.New(r)
	v.Name = "expense/ExpenseCreate"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Clientes"] = res

	v.Render(w)

	//tmpl.ExecuteTemplate(w, "ExpenseCreate", res)
}

func ExpenseStore(w http.ResponseWriter, r *http.Request) {

	// Check the request form METHOD
	if r.Method == "POST" {
		file, handler, err := r.FormFile("bill_url") //r.FormValue("recepit")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./expenses/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		expense := model.Expenses{}
		// Get the values from Form
		expense.Types = r.FormValue("types")
		expense.Status = r.FormValue("status")
		expense.Bill_url = "./expenses/" + handler.Filename
		expense.Recepit = "./expenses/" + handler.Filename
		expense.Payment, _ = strconv.ParseFloat(r.FormValue("payment"), 64)
		expense.Client_id, _ = strconv.Atoi(r.FormValue("cliente"))
		expense.Datedue = r.FormValue("datedue")
		expense.Datepaid = r.FormValue("datepaid")
		expense.Expensename = r.FormValue("expensename")
		model.ExpenseStore(expense)

	}

	// Redirect to HOME
	http.Redirect(w, r, "/expense", 307)
}
