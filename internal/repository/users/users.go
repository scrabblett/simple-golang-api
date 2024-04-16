package users

import (
	"context"
	"database/sql"
	"errors"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/repository/users/model"
	"time"
)

//go:generate mockery --name UsersRepo
type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (repo *UsersRepo) GetUserCredentials(ctx context.Context, login string) (model.UserCredentials, error) {
	var credentials model.UserCredentials
	tx, err := repo.db.Begin()

	if err != nil {
		return model.UserCredentials{}, err
	}

	script := "SELECT login, password, salt, user_id FROM users_info WHERE login = $1"

	err = tx.QueryRowContext(
		ctx, script, login,
	).Scan(&credentials.Login, &credentials.Password, &credentials.Salt, &credentials.UserId)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserCredentials{}, domain.ErrInvalidCredentials
		}

		return model.UserCredentials{}, err
	}

	err = tx.Commit()

	if err != nil {
		return model.UserCredentials{}, err
	}

	return credentials, nil
}

func (repo *UsersRepo) SaveUserCredentials(ctx context.Context, userInfo *model.SignUpUser) (int64, error) {
	tx, err := repo.db.Begin()

	if err != nil {
		return 0, err
	}

	script := `
		INSERT INTO library_user (last_name, first_name, patronymic, birth_date, created_at, edited_at, removed_at) 
		VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp, null)
		RETURNING id
	`

	var userId int64

	err = tx.QueryRowContext(ctx, script, userInfo.LastName, userInfo.FirstName, userInfo.Patronymic, userInfo.BirthDate).Scan(&userId)

	if err != nil {
		return 0, err
	}

	credsScript := `
		INSERT INTO users_info (login, password, salt, user_id) VALUES ($1, $2, $3, $4)
    `

	res, err := tx.ExecContext(ctx, credsScript, userInfo.Login, userInfo.Password, userInfo.Salt, userId)

	if err != nil {
		return 0, err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return 0, err
	}

	err = tx.Commit()

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (repo *UsersRepo) SaveJWTToken(ctx context.Context, userId int64, jwt string) error {
	tx, err := repo.db.Begin()

	if err != nil {
		return err
	}

	script := `INSERT INTO users_tokens (token, expired_at, user_id) VALUES ($1, $2, $3)`

	res, err := tx.ExecContext(ctx, script, jwt, time.Now().AddDate(0, 0, 7), userId)

	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()

	if rows == 0 {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (repo *UsersRepo) GetJWTToken(ctx context.Context, userId int64) (string, error) {
	var jwtToken string

	script := `SELECT token FROM users_tokens WHERE user_id = $1 AND expired_at > current_timestamp`

	err := repo.db.QueryRowContext(ctx, script, userId).Scan(&jwtToken)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.ErrTokenExpired
		}

		return "", err
	}

	return jwtToken, nil
}
