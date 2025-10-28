package models

import (
	"time"
)

type TransactionGroup struct {
	Id         int64     `json:"id"`
	EstateId   int64     `json:"estate_id"`
	SettingId  *int64    `json:"setting_id"`
	Direction  bool      `json:"direction"`
	Regularity bool      `json:"regularity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewTransactionGroup(id, estateId int64, settingId *int64, direction, regularity bool) *TransactionGroup {
	return &TransactionGroup{
		Id:         id,
		EstateId:   estateId,
		SettingId:  settingId,
		Direction:  direction,
		Regularity: regularity,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

type CreateTransactionGroup struct {
	EstateId   int64  `json:"estate_id", db:"estate_id"`
	SettingId  *int64 `json:"setting_id", db:"setting_id"`
	Direction  bool   `json:"direction", db:"direction"`
	Regularity bool   `json:"regularity", db:"regularity"`
}

type UpdateTransactionGroup struct {
	Id         int64  `json:"id", db:"id"`
	EstateId   int64  `json:"estate_id", db:"estate_id"`
	SettingId  *int64 `json:"setting_id", db:"setting_id"`
	Direction  bool   `json:"direction", db:"direction"`
	Regularity bool   `json:"regularity", db:"regularity"`
}
