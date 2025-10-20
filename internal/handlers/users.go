package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"
    "time"

    "github.com/go-chi/chi/v5"
    "github.com/golang-jwt/jwt/v5"


    "github.com/Parovozzzik/real-estate-portfolio/internal/config"
    "github.com/Parovozzzik/real-estate-portfolio/internal/logging"
    "github.com/Parovozzzik/real-estate-portfolio/internal/models"
    "github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
}

type AuthResponse struct {
    Token string `json:"token"`
    User  *models.User `json:"user"`
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

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.Id,
        "email":   user.Email,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
        "iat":     time.Now().Unix(),
    })

    cfg := config.GetConfig()
    tokenString, err := token.SignedString([]byte(cfg.JwtSecret))
    if err != nil {
        http.Error(w, `{"error": "Failed to generate token"}`, http.StatusInternalServerError)
        return
    }

    response := AuthResponse{
        Token: tokenString,
        User:  user,
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

	registration.Password = string(hashedPassword)
    newUserId, err := h.userRepository.CreateUser(registration)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, `{"message": ` + err.Error() + `}`)
        return
    }

    w.WriteHeader(http.StatusCreated)
    strNewUserId := strconv.FormatInt(newUserId, 10)
    fmt.Fprintf(w, `{"id": ` + strNewUserId + `}`)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    user, err := h.userRepository.GetUserById(1)
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

    id := chi.URLParam(r, "id")

    jsonData, err := h.userRepository.GetUserEstates(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Write(jsonData)
}