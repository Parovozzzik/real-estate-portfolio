package models

import (
	"time"
)

type TransactionType struct {
	Id         int64     `json:"id"`
	Name       string    `json:"name"`
	Direction  bool      `json:"direction"`
	Regularity bool      `json:"regularity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewTransactionType(id int64, name string, direction, regularity bool) *TransactionType {
	return &TransactionType{
		Id:         id,
		Name:       name,
		Direction:  direction,
		Regularity: regularity,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type CreateTransactionType struct {
	Name       string `json:"name", db:"name"`
	Direction  bool   `json:"direction", db:"direction"`
	Regularity bool   `json:"regularity", db:"regularity"`
}

type UpdateTransactionType struct {
	Id         int64  `json:"id", db:"id"`
	Name       string `json:"name", db:"name"`
	Direction  *bool  `json:"direction", db:"direction"`
	Regularity *bool  `json:"regularity", db:"regularity"`
}
