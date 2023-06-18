package models

import "gorm.io/gorm"

type CoinType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
