package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Name      *string   `json:"name"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone"`
	Password  string    `json:"-"` // Omit from JSON output
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(id int64, username string, email string, password string) *User {
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
	Email    string `json:"email", db:"email"`
	Password string `json:"password", db:"password"`
}

type Login struct {
	Email    string `json:"email", db:"email"`
	Password string `json:"password", db:"password"`
}

type UpdateUser struct {
	Id    int64   `json:"id", db:"id"`
	Name  *string `json:"name", db:"name"`
	Email *string `json:"email", db:"email"`
	Phone *string `json:"phone", db:"phone"`
}
