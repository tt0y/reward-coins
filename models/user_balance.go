package models

type UserBalance struct {
	UserId     int `json:"user_id"`
	CoinTypeId int `json:"coin_type_id"`
	Amount     int `json:"amount"`
}
