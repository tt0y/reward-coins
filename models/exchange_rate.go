package models

type ExchangeRate struct {
	CoinTypeIdFrom int    `json:"coin_type_id_from"`
	CoinTypeIdTo   int    `json:"coin_type_id_to"`
	Rate           int    `json:"rate"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}
