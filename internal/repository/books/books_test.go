package books

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-golang-api/internal/repository/books/model"
	"simple-golang-api/internal/repository/mocks"
	"testing"
)

func TestInsertBook(t *testing.T) {
	book := model.Book{
		Title:          "Book Title",
		Description:    "Book Description",
		AgeGroup:       18,
		PublishingDate: "01-01-2000",
	}

	t.Run("success", func(t *testing.T) {
		mockedRepo := createMockedRepository()

		mockedRepo.On(
			"InsertBook", mock.Anything, mock.AnythingOfType("*model.Book"),
		).Return(nil)

		err := mockedRepo.InsertBook(context.TODO(), &book)

		assert.Nil(t, err)
		mockedRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockedRepo := createMockedRepository()

		mockedRepo.On(
			"InsertBook", mock.Anything, mock.AnythingOfType("*model.Book"),
		).Return(errors.New("error"))

		err := mockedRepo.InsertBook(context.TODO(), &book)

		assert.NotNil(t, err)
	})
}

func createMockedRepository() *mocks.BooksRepo {
	return new(mocks.BooksRepo)
}
