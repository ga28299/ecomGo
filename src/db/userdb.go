package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/ga28299/ecomGo/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := pgCon.Get(user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		// Handle other errors
		return nil, err
	}
	return user, nil
}

func GetUserByPhone(phone string) (*models.User, error) {
	user := &models.User{}
	err := pgCon.Get(user, "SELECT * FROM users WHERE phone=$1", phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		// Handle other errors
		return nil, err
	}
	return user, nil
}

func CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		INSERT INTO users (id, first_name, last_name, email, password, phone, token, refresh_token, created_at, updated_at, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		`
	_, err := pgCon.ExecContext(ctx, query, user.ID, *user.FirstName, *user.LastName, *user.Email, *user.Password, *user.Phone, *user.Token, *user.RefreshToken, user.CreatedAt, user.UpdatedAt, user.UserID)
	if err != nil {
		log.Println("Error in CreateUser", err)
		return err
	}

	return nil

}

func UpdateTokensForID(userID string, token string, refreshToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `
		UPDATE users SET token=$1, refresh_token=$2 WHERE user_id=$3
		`
	_, err := pgCon.ExecContext(ctx, query, token, refreshToken, userID)
	if err != nil {
		log.Println("Error in UpdateTokens", err)
		return err
	}

	return nil
}
