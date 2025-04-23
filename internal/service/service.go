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

//go:generate mockery
type BooksService interface {
	CreateNewBook(ctx context.Context, book *domain.Book) error
	GetBookById(ctx context.Context, id int64) (*domain.Book, error)
	UpdateBookById(ctx context.Context, id int64, book *domain.Book) error
	DeleteBookById(ctx context.Context, id int64) error
}

//go:generate mockery
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
