package features

type UpdateCardStockParams struct {
	Id        string
	Finish    string
	Language  string
	Stock     int64
	HasVendor bool
}
