package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"
    "strconv"

    "github.com/Parovozzzik/real-estate-portfolio/internal/logging"
    "github.com/Parovozzzik/real-estate-portfolio/internal/repositories"
    "github.com/Parovozzzik/real-estate-portfolio/internal/models"
)

type UserHandler struct {
	userRepository *repositories.UserRepository
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

    jsonData, err := json.Marshal(user)
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

    jsonData, err := h.userRepository.GetUsers()
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Write(jsonData)
}