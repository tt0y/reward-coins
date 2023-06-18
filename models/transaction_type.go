package models

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}
