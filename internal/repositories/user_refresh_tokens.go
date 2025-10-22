package repositories

import (
	"database/sql"
	"github.com/Parovozzzik/real-estate-portfolio/internal/models"
	"time"
)

type UserRefreshTokenRepository struct {
	db *sql.DB
}

func NewUserRefreshTokenRepository(db *sql.DB) *UserRefreshTokenRepository {
	return &UserRefreshTokenRepository{db: db}
}

func (u *UserRefreshTokenRepository) UpsertRefreshToken(createUserRefreshToken *models.CreateUserRefreshToken) (int64, error) {
	id, err := u.getRefreshTokenIdByUserId(createUserRefreshToken.UserId)
	if err == nil && id > 0 {
		err = u.UpdateUserRefreshTokenById(createUserRefreshToken.Token, createUserRefreshToken.ExpiresAt, id)
		return 0, err
	}

	return u.CreateRefreshToken(createUserRefreshToken)
}

func (u *UserRefreshTokenRepository) CreateRefreshToken(createUserRefreshToken *models.CreateUserRefreshToken) (int64, error) {
	result, err := u.db.Exec(
		"INSERT INTO real_estate_portfolio.rep_user_refresh_tokens (token, user_id, expires_at) VALUES (?, ?, ?)",
		createUserRefreshToken.Token,
		createUserRefreshToken.UserId,
		time.Unix(createUserRefreshToken.ExpiresAt, 0))
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (u *UserRefreshTokenRepository) GetUserIdByRefreshToken(refreshToken string) (int64, error) {
	query := "SELECT user_id FROM real_estate_portfolio.rep_user_refresh_tokens WHERE token = ? AND expires_at > NOW()"

	row, err := u.db.Query(query, refreshToken)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var userId int64 = 0
	row.Next()
	err = row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *UserRefreshTokenRepository) getRefreshTokenIdByUserId(userId int64) (int64, error) {
	query := "SELECT id as count FROM real_estate_portfolio.rep_user_refresh_tokens WHERE user_id = ?"

	row, err := u.db.Query(query, userId)
	defer row.Close()
	if err != nil {
		return 0, err
	}

	var id int64 = 0
	row.Next()
	err = row.Scan(&id)
	return id, err
}

func (u *UserRefreshTokenRepository) UpdateUserRefreshToken(oldRefreshToken, newRefreshToken string, expiresAt int64) error {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_user_refresh_tokens SET token = ?, expires_at = ? WHERE token = ?",
		newRefreshToken,
		time.Unix(expiresAt, 0),
		oldRefreshToken)
	return err
}

func (u *UserRefreshTokenRepository) UpdateUserRefreshTokenById(newRefreshToken string, expiresAt, id int64) error {
	_, err := u.db.Exec(
		"UPDATE real_estate_portfolio.rep_user_refresh_tokens SET token = ?, expires_at = ? WHERE id = ?",
		newRefreshToken,
		time.Unix(expiresAt, 0),
		id)
	return err
}
