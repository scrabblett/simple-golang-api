package converter

import (
	"simple-golang-api/internal/domain"
	repoModel "simple-golang-api/internal/repository/books/model"
)

func ToBookFromService(book *domain.Book) *repoModel.Book {
	return &repoModel.Book{
		Id:             book.Id,
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}

func ToBookFromRepo(book *repoModel.Book) *domain.Book {
	return &domain.Book{
		Id:             book.Id,
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}
