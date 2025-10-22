package models

import (
	"time"
)

type Estate struct {
	Id           int64     `json:"id"`
	EstateTypeId int64     `json:"estate_type_id"`
	Name         string    `json:"name"`
	UserId       int64     `json:"user_id"`
	Active       int       `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewEstate(id int64, name string, estateTypeId int64, userId int64, active int) *Estate {
	return &Estate{
		Id:           id,
		Name:         name,
		EstateTypeId: estateTypeId,
		UserId:       userId,
		Active:       active,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

type CreateEstate struct {
	Name         string `json:"name", db:"name"`
	EstateTypeId int64  `json:"estate_type_id", db:"estate_type_id"`
	UserId       int64  `json:"user_id", db:"user_id"`
}

type UpdateEstate struct {
	Id           int64  `json:"id"`
	Name         string `json:"name", db:"name"`
	EstateTypeId int64  `json:"estate_type_id", db:"estate_type_id"`
}

type FullEstate struct {
	Id             int64     `json:"id"`
	EstateTypeId   int64     `json:"estate_type_id"`
	EstateTypeName string    `json:"estate_type_name"`
	EstateTypeIcon string    `json:"estate_type_icon"`
	Name           string    `json:"name"`
	UserId         int64     `json:"user_id"`
	Active         int       `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
