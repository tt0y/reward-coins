package models

import "gorm.io/gorm"

type ExchangeRate struct {
	gorm.Model
	CoinTypeIdFrom int `json:"coin_type_id_from"`
	CoinTypeIdTo   int `json:"coin_type_id_to"`
	Rate           int `json:"rate"`
}
