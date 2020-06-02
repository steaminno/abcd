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

func Product(w http.ResponseWriter, r *http.Request) {

	sess := session.Instance(r)
	result := model.Product()
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "product/Product"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Vars["Produtos"] = result
	v.Render(w)

}

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)
	//fmt.Printf("MIME Header: %+v\n", result)
	v := view.New(r)
	v.Name = "product/ProductCreate"
	v.Vars["token"] = csrfbanana.Token(w, r, sess)
	v.Render(w)

}

func ProductStore(w http.ResponseWriter, r *http.Request) {

	produtos := model.Produtos{}

	if r.Method == "POST" {
		file, handler, err := r.FormFile("pimg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./products/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
		// Get the values from Form
		produtos.Pname = r.FormValue("pname")
		produtos.Pdescription = r.FormValue("pdescription")
		produtos.Valorunidade, _ = strconv.ParseFloat(r.FormValue("valorunidade"), 64)
		produtos.Product_cost, _ = strconv.ParseFloat(r.FormValue("product_cost"), 64)
		produtos.Pdatafabi = r.FormValue("pdatafabi")
		produtos.Pimg = "./products/" + handler.Filename
		produtos.Status = r.FormValue("status")
		produtos.Obs = r.FormValue("obs")

		model.ProductStore(produtos)

		// Prepare a SQL INSERT and check for errors
	}

	// Close database connection

	// Redirect to
	// sess := session.Instance(r)
	// //fmt.Printf("MIME Header: %+v\n", result)
	// v := view.New(r)
	// v.Name = "product/Product"
	// v.Vars["token"] = csrfbanana.Token(w, r, sess)
	// v.Render(w)

	http.Redirect(w, r, "/product", 307)
}
