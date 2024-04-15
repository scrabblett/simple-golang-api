package repository

import (
	"awesomeProject/internal/repository/books"
	bookModels "awesomeProject/internal/repository/books/model"
	"awesomeProject/internal/repository/users"
	userModels "awesomeProject/internal/repository/users/model"
	"context"
	"database/sql"
)

//go:generate mockery --name UsersRepo
type UsersRepo interface {
	GetUserCredentials(ctx context.Context, login string) (userModels.UserCredentials, error)
	SaveUserCredentials(ctx context.Context, userInfo *userModels.SignUpUser) (int64, error)
	SaveJWTToken(ctx context.Context, userId int64, jwt string) error
	GetJWTToken(ctx context.Context, userId int64) (string, error)
}

//go:generate mockery --name BooksRepo
type BooksRepo interface {
	InsertBook(ctx context.Context, book *bookModels.Book) error
}

type Repositories struct {
	Books BooksRepo
	Users UsersRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Books: books.NewBookRepo(db),
		Users: users.NewUsersRepo(db),
	}
}
