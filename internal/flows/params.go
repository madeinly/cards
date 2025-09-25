package flows

type RegisterCardParams struct {
	ID         string `json:"scryfall_id"`
	Vendor     string `json:"vendor"`
	Language   string `json:"language"`
	Finish     string `json:"finish"`
	Stock      int64  `json:"stock"`
	Visibility string `json:"visibility"`
}

type GetDashboardCardsParams struct {
	SetCode  string
	CardName string
	Page     string
	Limit    string
}
