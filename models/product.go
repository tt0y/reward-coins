package models

type Product struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Cost       int    `json:"cost"`
	CoinTypeId int    `json:"coin_type_id"`
	Images     string `json:"images"`
	Active     bool   `json:"active"`
}
