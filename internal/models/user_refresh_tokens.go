package models

import (
	"time"
)

type UserRefreshToken struct {
	Id        int64     `json:"id"`
	Token     string    `json:"token"`
	UserId    int64     `json:"user_id"`
	ExpiresAt int64     `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserRefreshToken(id int64, token string, userId int64, expiresAt int64) *UserRefreshToken {
	return &UserRefreshToken{
		Id:        id,
		Token:     token,
		UserId:    userId,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type CreateUserRefreshToken struct {
	Token     string `json:"token", db:"token"`
	UserId    int64  `json:"user_id", db:"user_id"`
	ExpiresAt int64  `json:"expires_at", db:"expires_at"`
}
