package repository

import (
	"context"
	"database/sql"
	"simple-golang-api/internal/repository/books"
	bookModels "simple-golang-api/internal/repository/books/model"
	"simple-golang-api/internal/repository/users"
	userModels "simple-golang-api/internal/repository/users/model"
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
	GetBook(ctx context.Context, id int64) (*bookModels.Book, error)
	DeleteBook(ctx context.Context, id int64) error
	UpdateBook(ctx context.Context, id int64, book *bookModels.Book) error
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
