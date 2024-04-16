package books

import (
	"context"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/repository"
	"simple-golang-api/internal/repository/books/converter"
)

type BookService struct {
	repo repository.BooksRepo
}

func NewBooksService(repo repository.BooksRepo) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) CreateNewBook(ctx context.Context, book *domain.Book) error {
	repoBook := converter.ToBookFromService(book)

	err := b.repo.InsertBook(ctx, repoBook)

	if err != nil {
		return err
	}

	book.Id = repoBook.Id

	return nil
}
