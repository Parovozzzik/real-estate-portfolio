package handlers

import (
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

type UserRefreshTokenHandler struct {
	userRefreshTokenRepository *repositories.UserRefreshTokenRepository
}

func NewUserRefreshTokenHandler(userRefreshTokenRepository *repositories.UserRefreshTokenRepository) *UserRefreshTokenHandler {
	return &UserRefreshTokenHandler{userRefreshTokenRepository: userRefreshTokenRepository}
}
