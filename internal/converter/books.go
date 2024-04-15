package converter

import (
	"awesomeProject/internal/domain"
	desc "awesomeProject/pkg/v1/book"
)

func ToBookFromDesc(book *desc.Book) *domain.Book {
	return &domain.Book{
		Title:          book.Title,
		Description:    book.Description,
		PublishingDate: book.PublishingDate,
		AgeGroup:       book.AgeGroup,
	}
}
