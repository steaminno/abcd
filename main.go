package main

import (
	// Database SQL package to perform queries
	"encoding/json"
	"log"      // Display messages to console
	"net/http" // Manage URL
	"os"

	// "os"
	// "runtime"

	"github.com/SmartHobbyjd/config" // Database, serve port, email, uploaded files

	"github.com/SmartHobbyjd/gerenciador/shared/database" // Database, serve port, email, uploaded files
	"github.com/SmartHobbyjd/gerenciador/shared/jsonconfig"
	"github.com/SmartHobbyjd/gerenciador/shared/recaptcha"
	"github.com/SmartHobbyjd/gerenciador/shared/server"
	"github.com/SmartHobbyjd/gerenciador/shared/session"
	"github.com/SmartHobbyjd/gerenciador/shared/view"
	"github.com/SmartHobbyjd/gerenciador/shared/view/plugin"

	//
	//
	myroute "github.com/SmartHobbyjd/gerenciador/route"
	// "./middleware"  //
	// "encoding/json"
	/*
		"app/route"
		"app/shared/database"
		"app/shared/email"
		"app/shared/jsonconfig"
		"app/shared/recaptcha"
		"app/shared/server"
		"app/shared/session"
		"app/shared/view"
		"app/shared/view/plugin"
	*/)

func main() {

	//r := httprouter.New()

	// Show on console the application stated
	log.Print("Server startd on ", config.BaseUrl)

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", gconfig)
	database.Connect(gconfig.Database)
	session.Configure(gconfig.Session)
	view.Configure(gconfig.View)
	recaptcha.Configure(gconfig.Recaptcha)

	view.LoadTemplates(gconfig.Template.Root, gconfig.Template.Children)
	view.LoadPlugins(
		plugin.TagHelper(gconfig.View),
		plugin.NoEscape(),
		plugin.PrettyTime(),
		recaptcha.Plugin())

	http.Handle("/products/", http.StripPrefix("/products", http.FileServer(http.Dir("products/"))))

	http.Handle("/expenses/", http.StripPrefix("/expenses", http.FileServer(http.Dir("expenses/"))))

	http.Handle("/cheques/", http.StripPrefix("/cheques", http.FileServer(http.Dir("cheques/"))))

	//routes := myroute.LoadHTTP()

	server.Run(myroute.LoadHTTP(), myroute.LoadHTTPS(), gconfig.Server)

	// err := http.ListenAndServe(":"+config.Port, routes) // setting listening port
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}

// func oldmain() {

// 	// Show on console the application stated
// 	log.Print("Server startd on ", config.BaseUrl)
// 	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", gconfig)
// 	database.Connect(gconfig.Database)
// 	session.Configure(gconfig.Session)
// 	view.Configure(gconfig.View)
// 	recaptcha.Configure(gconfig.Recaptcha)

// 	view.LoadTemplates(gconfig.Template.Root, gconfig.Template.Children)
// 	view.LoadPlugins(
// 		plugin.TagHelper(gconfig.View),
// 		plugin.NoEscape(),
// 		plugin.PrettyTime(),
// 		recaptcha.Plugin())

// 	// Manage templates

// 	/* Default Routing */
// 	IndexHandler := http.HandlerFunc(appController.Index)
// 	ShowHandler := http.HandlerFunc(appController.Show)
// 	Showamount := http.HandlerFunc(controller.Showamount)
// 	NewHandler := http.HandlerFunc(appController.New)
// 	EditHandler := http.HandlerFunc(appController.Edit)
// 	// Maverage
// 	MaverageHandler := http.HandlerFunc(appController.Maverage)
// 	MaverageShowHandler := http.HandlerFunc(appController.MaverageShow)
// 	MaverageUpdateHandler := http.HandlerFunc(appController.MaverageUpdate)
// 	MaverageEditHandler := http.HandlerFunc(appController.MaverageEdit)
// 	MaverageCreateHandler := http.HandlerFunc(appController.MaverageCreate)
// 	MaverageDeleteHandler := http.HandlerFunc(appController.MaverageDelete)
// 	// Benefics
// 	BeneficHandler := http.HandlerFunc(appController.Benefic)
// 	BeneficCreateHandler := http.HandlerFunc(appController.BeneficCreate)
// 	ClientHandler := http.HandlerFunc(appController.Client)
// 	ClientCreateHandler := http.HandlerFunc(appController.ClientCreate)
// 	ClientStoreHandler := http.HandlerFunc(appController.ClientStore)
// 	ChequeHandler := http.HandlerFunc(appController.Cheque)
// 	ChequeCreateHandler := http.HandlerFunc(appController.ChequeCreate)
// 	ChequeStoreHandler := http.HandlerFunc(appController.ChequeStore)
// 	ProductHandler := http.HandlerFunc(appController.Product)
// 	ProductCreateHandler := http.HandlerFunc(appController.ProductCreate)
// 	RegistrationHandler := http.HandlerFunc(appController.Registration)
// 	RegistrationCreateHandler := http.HandlerFunc(appController.RegistrationCreate)
// 	RegistrationStoreHandler := http.HandlerFunc(appController.RegistrationStore)
// 	InsertHandler := http.HandlerFunc(appController.Insert)
// 	UpdateHandler := http.HandlerFunc(appController.Update)
// 	DeleteHandler := http.HandlerFunc(appController.Delete)
// 	MaverageStoreHandler := http.HandlerFunc(appController.MaverageStore)
// 	BeneficStoreHandler := http.HandlerFunc(appController.BeneficStore)
// 	ProductStoreHandler := http.HandlerFunc(appController.ProductStore)
// 	ExpenseHandler := http.HandlerFunc(appController.Expense)
// 	ExpenseCreateHandler := http.HandlerFunc(appController.ExpenseCreate)
// 	ExpenseStoreHandler := http.HandlerFunc(appController.ExpenseStore)
// 	//HomeHandler := http.HandlerFunc(controller.Home)
// 	//LoginHandler := http.HandlerFunc(controller.Login)
// 	//LogoutHandler := http.HandlerFunc(controller.Logout)
// 	//RegisterHandler := http.HandlerFunc(controller.Register)

// 	// URL management

// 	// /*********product**************/
// 	// http.Handle("/product/store", middleware.NoCache(ProductStoreHandler))
// 	http.Handle("/", IndexHandler)         // INDEX :: Show all registers
// 	http.Handle("/show", ShowHandler)      // SHOW  :: Show only one register
// 	http.Handle("/showamount", Showamount) // SHOW  :: Show total Valortotal per status
// 	http.Handle("/new", NewHandler)        // NEW   :: Form to create new register
// 	http.Handle("/edit", EditHandler)      // EDIT  :: Form to edit register

// 	/*******client************/
// 	http.Handle("/client", ClientHandler)              // CLIENT :: Show all Clinets
// 	http.Handle("/client/create", ClientCreateHandler) //CLIENTCREATE :: Client Register Form
// 	/*******cheque************/
// 	http.Handle("/cheque", ChequeHandler)              // CHEQUE :: Show all cheques
// 	http.Handle("/cheque/create", ChequeCreateHandler) //CHEQUECREATE :: Cheque Register Form
// 	/*******maverage************/
// 	http.Handle("/maverage", MaverageHandler)              // Maverage :: Show all Maverages
// 	http.Handle("/maverageshow", MaverageShowHandler)      // MaverageShow :: Show by ID
// 	http.Handle("/maverageupdate", MaverageUpdateHandler)  // Maverageupdate :: update by ID
// 	http.Handle("/maverageedit", MaverageEditHandler)      // MaverageEdit :: Edit by ID
// 	http.Handle("/maveragedelete", MaverageDeleteHandler)  // MaveragesDelete:: Delete By ID
// 	http.Handle("/maverage/create", MaverageCreateHandler) //CLIENTCREATE :: Maverage Register Form
// 	/*******benefic************/
// 	http.Handle("/benefic", BeneficHandler)              // Benefic :: Show all Benefics
// 	http.Handle("/benefic/create", BeneficCreateHandler) //CLIENTCREATE :: Benefic Register Form
// 	/******product************/
// 	http.Handle("/product", ProductHandler)              // PRODUCT :: Show all Products
// 	http.Handle("/product/create", ProductCreateHandler) //PRODUCTCREATE :: Product Register Form
// 	/*******registration************/
// 	http.Handle("/registration", RegistrationHandler)              // Registration :: Show all Users
// 	http.Handle("/registration/create", RegistrationCreateHandler) //RegistrationCREATE :: Users Register Form

// 	// Manage actions
// 	/*********main**************/
// 	http.Handle("/insert", InsertHandler) // INSERT :: New register
// 	http.Handle("/update", UpdateHandler) // UPDATE :: Update register
// 	http.Handle("/delete", DeleteHandler) // DELETE :: Destroy register
// 	/*********maverage**************/
// 	http.Handle("/maverage/store", MaverageStoreHandler)
// 	/*********benefic**************/
// 	http.Handle("/benefic/store", BeneficStoreHandler)
// 	/*********client**************/
// 	http.Handle("/client/store", ClientStoreHandler)
// 	/*********cheque**************/
// 	http.Handle("/cheque/store", ChequeStoreHandler)
// 	/*********product**************/
// 	http.Handle("/product/store", ProductStoreHandler)
// 	http.HandleFunc("/uploads/", serveResource)
// 	/*********registration**************/
// 	http.Handle("/registration/store", RegistrationStoreHandler)

// 	/***************expense**************/

// 	http.Handle("/expense", ExpenseHandler)
// 	http.Handle("/expense/create", ExpenseCreateHandler)
// 	http.Handle("/expense/store", ExpenseStoreHandler)

// 	/***************uploaded files folders**************/
// 	http.Handle("/products/", http.StripPrefix("/products", http.FileServer(http.Dir("products/"))))
// 	http.Handle("/expenses/", http.StripPrefix("/expenses", http.FileServer(http.Dir("expenses/"))))
// 	http.Handle("/cheques/", http.StripPrefix("/cheques", http.FileServer(http.Dir("cheques/"))))

// 	// Start the server on port 9000
// 	err := http.ListenAndServe(":"+config.Port, nil) // setting listening port
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "./public" + req.URL.Path
	http.ServeFile(w, req, path)
}

// config the settings variable
var gconfig = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Info `json:"Database"`
	// Email     email.SMTPInfo  `json:"Email"`
	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
