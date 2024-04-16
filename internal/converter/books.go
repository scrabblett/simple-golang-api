package converter

import (
	"simple-golang-api/internal/domain"
	desc "simple-golang-api/pkg/v1/book"
)

func ToBookFromDesc(book *desc.Book) *domain.Book {
	return &domain.Book{
		Id:             book.Id,
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}

func ToBookFromService(book *domain.Book) *desc.Book {
	return &desc.Book{
		Id:             book.Id,
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}
