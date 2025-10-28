package models

import (
	"time"
)

type Transaction struct {
	Id        int64     `json:"id"`
	GroupId   int64     `json:"group_id"`
	TypeId    int64     `json:"type_id"`
	Sum       float64   `json:"sum"`
	Date      time.Time `json:"date"`
	Comment   *string   `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTransaction(id, groupId, typeId int64, sum float64, date time.Time, comment *string) *Transaction {
	return &Transaction{
		Id:        id,
		GroupId:   groupId,
		TypeId:    typeId,
		Sum:       sum,
		Date:      date,
		Comment:   comment,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateTransaction struct {
	GroupId int64     `json:"group_id", db:"group_id"`
	TypeId  int64     `json:"type_id", db:"type_id"`
	Sum     float64   `json:"sum", db:"sum"`
	Date    time.Time `json:"date", db:"date"`
	Comment *string   `json:"comment", db:"comment"`
}

type UpdateTransaction struct {
	Id      int64     `json:"id", db:"id"`
	GroupId int64     `json:"group_id", db:"group_id"`
	TypeId  int64     `json:"type_id", db:"type_id"`
	Sum     float64   `json:"sum", db:"sum"`
	Date    time.Time `json:"date", db:"date"`
	Comment *string   `json:"comment", db:"comment"`
}
