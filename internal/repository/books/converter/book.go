package converter

import (
	"simple-golang-api/internal/domain"
	repoModel "simple-golang-api/internal/repository/books/model"
)

func ToBookFromService(book *domain.Book) *repoModel.Book {
	return &repoModel.Book{
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}
