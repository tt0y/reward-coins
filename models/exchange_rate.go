package models

type ExchangeRate struct {
	ID             uint `json:"id"`
	CoinTypeIdFrom int  `json:"coin_type_id_from"`
	CoinTypeIdTo   int  `json:"coin_type_id_to"`
	Rate           int  `json:"rate"`
}
