package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Omit from JSON output
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(id int, username string, email string, password string) *User {
	return &User{
		Id:        id,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Registration struct {
	Email string `json:"email", db:"email"`
	Password string `json:"password", db:"password"`
}

type Login struct {
	Email string `json:"email", db:"email"`
	Password string `json:"password", db:"password"`
}
