package models

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	CoinTypeId  int    `json:"coin_type_id"`
	Images      string `json:"images"`
	Active      bool   `json:"active"`
}
