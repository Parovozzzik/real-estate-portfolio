package models

import (
	"time"
)

type TransactionFrequency struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTransactionFrequency(id int64, name string) *TransactionFrequency {
	return &TransactionFrequency{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateTransactionFrequency struct {
	Name string `json:"name", db:"name"`
}

type UpdateTransactionFrequency struct {
	Id   int64  `json:"id", db:"id"`
	Name string `json:"name", db:"name"`
}
