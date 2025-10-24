package models

import (
	"time"
)

type TransactionType struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTransactionType(id int64, name string) *TransactionType {
	return &TransactionType{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateTransactionType struct {
	Name string `json:"name", db:"name"`
}

type UpdateTransactionType struct {
	Id   int64  `json:"id", db:"id"`
	Name string `json:"name", db:"name"`
}
