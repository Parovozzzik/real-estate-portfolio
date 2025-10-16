package models

import (
	"time"
)

type EstateType struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
    Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewEstateType(id int, name string, active int) *EstateType {
	return &EstateType{
		Id:        id,
		Name:      name,
		Active:    active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateEstateType struct {
	Name string `json:"name", db:"name"`
}
