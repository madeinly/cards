package flows

type Card struct {
	CardBase
	Language  string `json:"card_language"`
	Finish    string `json:"card_finish"`
	HasVendor bool   `json:"card_HasVendor"`
}

type CardBase struct {
	ID        string  `json:"card_id"`
	NameEN    string  `json:"card_nameEn"`
	NameES    string  `json:"card_nameES"`
	ImageURL  string  `json:"card_imageURL"`
	SetCode   string  `json:"card_setCode"`
	SetName   string  `json:"card_setName"`
	ManaValue int64   `json:"card_manaValue"`
	Number    string  `json:"card_number"`
	Rarity    string  `json:"card_rarity"`
	Colors    string  `json:"card_colors"`
	Types     string  `json:"card_types"`
	Price     float64 `json:"card_price"`
	Stock     int64   `json:"card_stock"`
}

type RegisterCardParams struct {
	ScryfallId string `json:"scryfall_id"`
	Vendor     string `json:"vendor"`
	Language   string `json:"language"`
	Finish     string `json:"finish"`
	Stock      string `json:"stock"`
	Visibility string `json:"visibility"`
}
