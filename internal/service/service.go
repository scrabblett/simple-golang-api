package service

import (
	"context"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/repository"
	"simple-golang-api/internal/service/books"
	"simple-golang-api/internal/service/users"
)

type Services struct {
	Books BooksService
	Users UsersService
}

type Deps struct {
	Repos *repository.Repositories
}

//go:generate mockery --name BooksService
type BooksService interface {
	CreateNewBook(ctx context.Context, book *domain.Book) error
}

//go:generate mockery --name UsersService
type UsersService interface {
	SignIn(ctx context.Context, credentials *domain.UserCredentials) (string, error)
	SignUp(ctx context.Context, userInfo *domain.SignUpUser) error
}

func NewServices(deps Deps) *Services {
	return &Services{
		Books: books.NewBooksService(deps.Repos.Books),
		Users: users.NewUsersService(deps.Repos.Users),
	}
}
