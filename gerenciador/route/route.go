package route

import (
	"net/http"

	"github.com/SmartHobbyjd/gerenciador/controller"
	"github.com/SmartHobbyjd/gerenciador/route/middleware/acl"

	hr "github.com/SmartHobbyjd/gerenciador/route/middleware/httprouterwrapper"
	"github.com/SmartHobbyjd/gerenciador/route/middleware/logrequest"
	"github.com/SmartHobbyjd/gerenciador/shared/session"

	"github.com/gorilla/context"
	"github.com/josephspurrier/csrfbanana"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	//"github.com/SmartHobbyjd/gerenciador/controller"
	appcontroller "github.com/SmartHobbyjd/gerenciador/controller"
)

// Load returns the routes and middleware
func Load() http.Handler {
	return middleware(routes())
}

// LoadHTTPS returns the HTTP routes and middleware
func LoadHTTPS() http.Handler {
	return middleware(routes())
}

// LoadHTTP returns the HTTPS routes and middleware
func LoadHTTP() http.Handler {
	return middleware(routes())

	// Uncomment this and comment out the line above to always redirect to HTTPS
	//return http.HandlerFunc(redirectToHTTPS)
}

// Optional method to make it easy to redirect from HTTP to HTTPS
func redirectToHTTPS(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "https://"+req.Host, http.StatusMovedPermanently)
}

// *****************************************************************************
// Routes
// *****************************************************************************

/* Default Routing */
//IndexHandler := http.HandlerFunc(controller.Index)
//ShowHandler := http.HandlerFunc(controller.Show)
//Showamount := http.HandlerFunc(controller.Showamount)
//NewHandler := http.HandlerFunc(controller.New)
//EditHandler := http.HandlerFunc(controller.Edit)
// Maverage
// MaverageHandler := http.HandlerFunc(controller.Maverage)
// MaverageShowHandler := http.HandlerFunc(controller.MaverageShow)
// MaverageUpdateHandler := http.HandlerFunc(controller.MaverageUpdate)
// MaverageEditHandler := http.HandlerFunc(controller.MaverageEdit)
// MaverageCreateHandler := http.HandlerFunc(controller.MaverageCreate)
// MaverageDeleteHandler := http.HandlerFunc(controller.MaverageDelete)
// // Benefics
// BeneficHandler := http.HandlerFunc(controller.Benefic)
// BeneficCreateHandler := http.HandlerFunc(controller.BeneficCreate)
// ClientHandler := http.HandlerFunc(controller.Client)
// ClientCreateHandler := http.HandlerFunc(controller.ClientCreate)
// ClientStoreHandler := http.HandlerFunc(controller.ClientStore)
// ChequeHandler := http.HandlerFunc(controller.Cheque)
// ChequeCreateHandler := http.HandlerFunc(controller.ChequeCreate)
// ChequeStoreHandler := http.HandlerFunc(controller.ChequeStore)
// ProductHandler := http.HandlerFunc(controller.Product)
// ProductCreateHandler := http.HandlerFunc(controller.ProductCreate)
// RegistrationHandler := http.HandlerFunc(controller.Registration)
// RegistrationCreateHandler := http.HandlerFunc(controller.RegistrationCreate)
// RegistrationStoreHandler := http.HandlerFunc(controller.RegistrationStore)
// InsertHandler := http.HandlerFunc(controller.Insert)
// UpdateHandler := http.HandlerFunc(controller.Update)
// DeleteHandler := http.HandlerFunc(controller.Delete)
// MaverageStoreHandler := http.HandlerFunc(controller.MaverageStore)
// BeneficStoreHandler := http.HandlerFunc(controller.BeneficStore)
// ProductStoreHandler := http.HandlerFunc(controller.ProductStore)
// ExpenseHandler := http.HandlerFunc(controller.Expense)
// ExpenseCreateHandler := http.HandlerFunc(controller.ExpenseCreate)
// ExpenseStoreHandler := http.HandlerFunc(controller.ExpenseStore)
//HomeHandler := http.HandlerFunc(controller.Home)
//LoginHandler := http.HandlerFunc(controller.Login)
//LogoutHandler := http.HandlerFunc(controller.Logout)
//RegisterHandler := http.HandlerFunc(controller.Register)

func routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.IndexGET)))

	// Login
	r.POST("/login", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.LoginPOST)))

	r.GET("/login", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.LoginGET)))

	r.GET("/logout", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.LogoutGET)))

	r.GET("/home", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Index)))

	//http.Handle("/", IndexHandler)
	r.GET("/show", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.Show))) // INDEX :: Show all registers
	//http.Handle("/show", ShowHandler)      // SHOW  :: Show only one register
	// r.GET("/showamount", hr.Handler(alice.
	// 	New().
	// 	ThenFunc(appcontroller.Showamount)))

	//http.Handle("/showamount", Showamount) // SHOW  :: Show total Valortotal per status

	r.GET("/new", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.New)))

	//http.Handle("/new", NewHandler)        // NEW   :: Form to create new register

	r.GET("/edit", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Edit)))

	/*******client************/
	// http.Handle("/client", ClientHandler)              // CLIENT :: Show all Clinets
	r.GET("/client", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Client)))

	r.POST("/client", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Client)))

	// http.Handle("/client/create", ClientCreateHandler) //CLIENTCREATE :: Client Register Form
	r.GET("/client/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ClientCreate)))

	// /*******cheque************/
	// http.Handle("/cheque", ChequeHandler)              // CHEQUE :: Show all cheques
	r.GET("/cheque", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Cheque)))

	r.POST("/cheque", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Cheque)))

	// http.Handle("/cheque/create", ChequeCreateHandler) //CHEQUECREATE :: Cheque Register Form
	r.GET("/cheque/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ChequeCreate)))

	// /*******maverage************/
	// MaverageHandler := http.HandlerFunc(controller.Maverage)
	// MaverageShowHandler := http.HandlerFunc(controller.MaverageShow)
	// MaverageUpdateHandler := http.HandlerFunc(controller.MaverageUpdate)
	// MaverageEditHandler := http.HandlerFunc(controller.MaverageEdit)
	// MaverageCreateHandler := http.HandlerFunc(controller.MaverageCreate)
	// MaverageDeleteHandler := http.HandlerFunc(controller.MaverageDelete)

	// http.Handle("/maverage", MaverageHandler)              // Maverage :: Show all Maverages
	r.GET("/maverage", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Maverage)))

	r.POST("/maverage", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Maverage)))

	// http.Handle("/maverageshow", MaverageShowHandler)      // MaverageShow :: Show by ID
	r.GET("/maverageshow", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageShow)))

	// http.Handle("/maverageupdate", MaverageUpdateHandler)  // Maverageupdate :: update by ID
	r.POST("/maverageupdate", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageUpdate)))

	// http.Handle("/maverageedit", MaverageEditHandler)      // MaverageEdit :: Edit by ID
	r.POST("/maverageedit", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageEdit)))

	r.GET("/maverageedit", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageEdit)))

	// http.Handle("/maveragedelete", MaverageDeleteHandler)  // MaveragesDelete:: Delete By ID
	r.GET("/maveragedelete", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageDelete)))

	// http.Handle("/maverage/create", MaverageCreateHandler) //CLIENTCREATE :: Maverage Register Form
	r.GET("/maverage/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageCreate)))

	// /*******benefic************/
	// http.Handle("/benefic", BeneficHandler)              // Benefic :: Show all Benefics
	r.GET("/benefic", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Benefic)))

	r.POST("/benefic", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Benefic)))
	// http.Handle("/benefic/create", BeneficCreateHandler) //CLIENTCREATE :: Benefic Register Form
	r.GET("/benefic/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.BeneficCreate)))

	// /******product************/
	// http.Handle("/product", ProductHandler)              // PRODUCT :: Show all Products
	r.GET("/product", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Product)))

	r.POST("/product", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Product)))
	// http.Handle("/product/create", ProductCreateHandler) //PRODUCTCREATE :: Product Register Form
	r.GET("/product/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ProductCreate)))

	// /*******registration************/
	// http.Handle("/registration", RegistrationHandler)              // Registration :: Show all Users
	r.GET("/registration", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Registration)))

	r.POST("/registration", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Registration)))

	// http.Handle("/registration/create", RegistrationCreateHandler) //RegistrationCREATE :: Users Register Form
	r.GET("/registration/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.RegistrationCreate)))

	// // Manage actions
	// /*********main**************/
	// http.Handle("/insert", InsertHandler) // INSERT :: New register
	r.POST("/insert", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Insert)))

	// http.Handle("/update", UpdateHandler) // UPDATE :: Update register
	r.POST("/update", hr.Handler(alice.
		New().
		ThenFunc(appcontroller.Update)))

	// http.Handle("/delete", DeleteHandler) // DELETE :: Destroy register
	r.POST("/delete", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Delete)))

	r.GET("/delete", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Delete)))

	// /*********maverage**************/
	// http.Handle("/maverage/store", MaverageStoreHandler)
	r.POST("/maverage/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.MaverageStore)))

	// /*********benefic**************/
	// http.Handle("/benefic/store", BeneficStoreHandler)
	r.POST("/benefic/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.BeneficStore)))

	// /*********client**************/
	// http.Handle("/client/store", ClientStoreHandler)
	r.POST("/client/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ClientStore)))

	// /*********cheque**************/
	// http.Handle("/cheque/store", ChequeStoreHandler)
	r.POST("/cheque/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ChequeStore)))

	// /*********product**************/
	// http.Handle("/product/store", ProductStoreHandler)
	r.POST("/product/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ProductStore)))

	// http.HandleFunc("/uploads/", serveResource)
	r.POST("/uploads/", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(serveResource)))

	// /*********registration**************/
	// http.Handle("/registration/store", RegistrationStoreHandler)
	r.POST("/registration/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.RegistrationStore)))

	// /***************expense**************/

	// http.Handle("/expense", ExpenseHandler)
	r.GET("/expense", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Expense)))

	r.POST("/expense", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.Expense)))

	// http.Handle("/expense/create", ExpenseCreateHandler)
	r.GET("/expense/create", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ExpenseCreate)))

	// http.Handle("/expense/store", ExpenseStoreHandler)
	r.POST("/expense/store", hr.Handler(alice.
		New(acl.DisallowAnon).
		ThenFunc(appcontroller.ExpenseStore)))

	r.ServeFiles("/products/*filepath", http.Dir("products/"))
	r.ServeFiles("/expenses/*filepath", http.Dir("expenses/"))
	r.ServeFiles("/cheques/*filepath", http.Dir("cheques/"))
	// /***************uploaded files folders**************/
	//http.Handle("/products/", http.StripPrefix("/products", http.Dir("products/")))

	// http.Handle("/expenses/", http.StripPrefix("/expenses", http.FileServer(http.Dir("expenses/"))))

	// http.Handle("/cheques/", http.StripPrefix("/cheques", http.FileServer(http.Dir("cheques/"))))

	return r
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "./public" + req.URL.Path
	http.ServeFile(w, req, path)
}

// *****************************************************************************
// Middleware
// *****************************************************************************

func middleware(h http.Handler) http.Handler {
	// Prevents CSRF and Double Submits
	cs := csrfbanana.New(h, session.Store, session.Name)
	cs.FailureHandler(http.HandlerFunc(controller.InvalidToken))
	cs.ClearAfterUsage(true)
	cs.ExcludeRegexPaths([]string{"/static(.*)"})
	csrfbanana.TokenLength = 32
	csrfbanana.TokenName = "token"
	csrfbanana.SingleToken = false
	h = cs

	// Log every request
	h = logrequest.Handler(h)

	// Clear handler for Gorilla Context
	h = context.ClearHandler(h)

	return h
}
