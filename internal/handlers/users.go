package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Parovozzzik/real-estate-portfolio/internal/database"
	"net/http"
	"strconv"
	"time"

	chi "github.com/go-chi/chi/v5"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/Parovozzzik/real-estate-portfolio/internal/config"
	"github.com/Parovozzzik/real-estate-portfolio/internal/logging"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	ExpiresAt    int64        `json:"expires_at"`
	User         *models.User `json:"user"`
}

func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
	return &UserHandler{userRepository: userRepository}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData, err := h.userRepository.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	login := &models.Login{}
	err := json.NewDecoder(r.Body).Decode(login)
	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.userRepository.LoginUser(login)
	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Генерируем access token (15 минут)
	accessToken, accessExp := generateAccessToken(user.Id, user.Email)

	// Генерируем refresh token (7 дней)
	refreshToken, refreshExp := generateRefreshToken(user.Id)

	if err := saveRefreshToken(user.Id, refreshToken, refreshExp); err != nil {
		logger.Println(err.Error())

		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp,
		User:         user,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		logger.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *UserHandler) RegistrationUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	registration := &models.Registration{}
	err := json.NewDecoder(r.Body).Decode(registration)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hashedPassword, err := models.HashPassword(registration.Password)
	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registration.Password = hashedPassword
	newUserId, err := h.userRepository.CreateUser(registration)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strNewUserId := strconv.FormatInt(newUserId, 10)
	fmt.Fprintf(w, `{"id": `+strNewUserId+`}`)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logging.Init()
	logger := logging.GetLogger()

	updateUser := &models.UpdateUser{}
	userId, err := strconv.ParseInt(chi.URLParam(r, "user-id"), 10, 64)
	if err != nil {
		panic(err)
	}
	updateUser.Id = userId

	err = json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err = h.userRepository.UpdateUser(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"message": `+err.Error()+`}`)
		return
	}

	w.WriteHeader(http.StatusCreated)
	strNewUserId := strconv.FormatInt(userId, 10)
	fmt.Fprintf(w, `{"id": `+strNewUserId+`}`)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.ParseInt(chi.URLParam(r, "user-id"), 10, 64)
	if err != nil {
		panic(err)
	}

	user, err := h.userRepository.GetUserById(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func (h *UserHandler) GetUserEstates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.ParseInt(chi.URLParam(r, "user-id"), 10, 64)
	if err != nil {
		panic(err)
	}
	jsonData, err := h.userRepository.GetUserEstates(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(jsonData)
}

func (h *UserHandler) GetUserEstate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId, err := strconv.ParseInt(chi.URLParam(r, "user-id"), 10, 64)
	if err != nil {
		panic(err)
	}
	estateId, err := strconv.ParseInt(chi.URLParam(r, "estate-id"), 10, 64)
	if err != nil {
		panic(err)
	}

	estate, err := h.userRepository.GetUserEstate(userId, estateId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(estate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
		return
	}

	// Валидируем refresh token
	userId, err := validateRefreshToken(req.RefreshToken)
	if err != nil {
		http.Error(w, `{"error": "Invalid refresh token"}`, http.StatusUnauthorized)
		return
	}

	// Получаем пользователя
	user, err := h.userRepository.GetUserById(userId)
	if err != nil {
		logging.Init()
		logger := logging.GetLogger()
		logger.Println(err.Error())
		http.Error(w, `{"error": "User not found"}`, http.StatusUnauthorized)
		return
	}

	// Генерируем новые токены
	accessToken, accessExp := generateAccessToken(user.Id, user.Email)
	refreshToken, refreshExp := generateRefreshToken(user.Id)

	// Обновляем refresh token в БД
	if err := updateRefreshToken(req.RefreshToken, refreshToken, refreshExp); err != nil {
		http.Error(w, `{"error": "Failed to update refresh token"}`, http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp,
		User:         user,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// Вспомогательные функции
func generateAccessToken(userID int64, email string) (string, int64) {
	exp := time.Now().Add(15 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     exp,
		"iat":     time.Now().Unix(),
		"type":    "access",
	})

	cfg := config.GetConfig()
	tokenString, _ := token.SignedString([]byte(cfg.JwtSecret))
	return tokenString, exp
}

func generateRefreshToken(userID int64) (string, int64) {
	exp := time.Now().Add(7 * 24 * time.Hour).Unix()

	// Используем UUID для refresh token вместо JWT
	token := uuid.New().String()

	return token, exp
}

// Сохранение refresh token в БД
func saveRefreshToken(userId int64, refreshToken string, expiresAt int64) error {
	db := database.GetDBInstance()
	userRefreshTokenRepository := repositories.NewUserRefreshTokenRepository(db)

	createUserRefreshToken := &models.CreateUserRefreshToken{
		Token:     refreshToken,
		UserId:    userId,
		ExpiresAt: expiresAt,
	}
	_, err := userRefreshTokenRepository.UpsertRefreshToken(createUserRefreshToken)
	return err
}

func validateRefreshToken(refreshToken string) (int64, error) {
	db := database.GetDBInstance()
	userRefreshTokenRepository := repositories.NewUserRefreshTokenRepository(db)

	return userRefreshTokenRepository.GetUserIdByRefreshToken(refreshToken)
}

func updateRefreshToken(oldToken, newToken string, expiresAt int64) error {
	db := database.GetDBInstance()
	userRefreshTokenRepository := repositories.NewUserRefreshTokenRepository(db)

	return userRefreshTokenRepository.UpdateUserRefreshToken(oldToken, newToken, expiresAt)
}
