package models

type Transaction struct {
	ID                uint   `json:"id"`
	Status            string `json:"status"`
	TransactionTypeId int    `json:"transaction_type_id"`
	CoinTypeId        int    `json:"coin_type_id"`
	Amount            int    `json:"amount"`
	UserIdFrom        int    `json:"user_id_from"`
	UserIdTo          int    `json:"user_id_to"`
	ProductId         int    `json:"product_id"`
}
