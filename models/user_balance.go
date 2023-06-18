package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserId     int `json:"user_id"`
	CoinTypeId int `json:"coin_type_id"`
	Amount     int `json:"amount"`
}
