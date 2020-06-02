package model

type Res struct {
	Orden      []Ordens
	Valortotal float64
	//Statuspaid        string
	//Statusopen        string
	//Statusparcial     string
	TotalOrden        int
	TotalClient       int
	TotalProduct      int
	TotalExpense      int
	TotalSaledpaid    int
	TotalSaledopen    int
	TotalSaledparcial int
	TotalSaled        float64
	TotalSellsPartial float64
	TotalSellsOpen    float64
}
