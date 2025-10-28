package models

import (
	"time"
)

type TransactionRepaymentPlan struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTransactionRepaymentPlan(id int64, name string) *TransactionRepaymentPlan {
	return &TransactionRepaymentPlan{
		Id:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateTransactionRepaymentPlan struct {
	Name string `json:"name", db:"name"`
}

type UpdateTransactionRepaymentPlan struct {
	Id   int64  `json:"id", db:"id"`
	Name string `json:"name", db:"name"`
}
