package models

import (
	"time"
)

type EstateType struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewEstateType(id int, name string, icon string, active int) *EstateType {
	return &EstateType{
		Id:        id,
		Name:      name,
		Icon:      icon,
		Active:    active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateEstateType struct {
	Name string `json:"name", db:"name"`
	Icon string `json:"icon", db:"icon"`
}
