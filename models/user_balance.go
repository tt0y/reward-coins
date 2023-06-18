package models

type UserBalance struct {
	ID         int `json:"id"`
	UserId     int `json:"user_id"`
	CoinTypeId int `json:"coin_type_id"`
	Amount     int `json:"amount"`
}
