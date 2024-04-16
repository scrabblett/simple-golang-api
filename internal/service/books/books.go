package books

import (
	"context"
	"go.uber.org/zap"
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
		zap.L().Error("failed to create book", zap.Error(err))

		return err
	}

	book.Id = repoBook.Id

	return nil
}

func (b *BookService) Get() {

}
