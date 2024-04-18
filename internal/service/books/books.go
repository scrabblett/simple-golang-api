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

func (b *BookService) GetBookById(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := b.repo.GetBook(ctx, id)

	if err != nil {
		return nil, err
	}

	serviceBook := converter.ToBookFromRepo(book)

	return serviceBook, nil
}

func (b *BookService) UpdateBookById(ctx context.Context, id int64, book *domain.Book) error {
	err := b.repo.UpdateBook(ctx, id, converter.ToBookFromService(book))

	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) DeleteBookById(ctx context.Context, id int64) error {
	err := b.repo.DeleteBook(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
