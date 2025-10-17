package models

import (
	"time"
)

type Estate struct {
	Id              int         `json:"id"`
	EstateTypeId    int         `json:"estate_type_id"`
	Name            string      `json:"name"`
	UserId          int         `json:"user_id"`
    Active          int         `json:"active"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}

func NewEstate(id int, name string, estateTypeId int, userId int, active int) *Estate {
	return &Estate{
		Id:             id,
		Name:           name,
		EstateTypeId:   estateTypeId,
		UserId:         userId,
        Active:         active,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

type CreateEstate struct {
	Name string `json:"name", db:"name"`
    EstateTypeId string `json:"estate_type_id", db:"estate_type_id"`
    UserId string `json:"user_id", db:"user_id"`
}
