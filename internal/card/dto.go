package card

type Card struct {
	ID        string   `json:"card_id"`
	Name      string   `json:"card_name"`
	Sku       string   `json:"card_sku"`
	ImageURL  string   `json:"card_imageURL"`
	SetCode   string   `json:"card_setCode"`
	SetName   string   `json:"card_setName"`
	ManaValue int64    `json:"card_manaValue"`
	Rarity    string   `json:"card_rarity"`
	Colors    []string `json:"card_colors"`
	Types     []string `json:"card_types"`
	Price     float64  `json:"card_price"`
}
