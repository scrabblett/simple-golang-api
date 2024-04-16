package books

import (
	"awesomeProject/internal/repository/books/model"
	"awesomeProject/internal/repository/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestInsertBook(t *testing.T) {
	book := model.Book{
		Title:          "Book Title",
		Description:    "Book Description",
		AgeGroup:       18,
		PublishingDate: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockedRepo := createMockedRepository()

		mockedRepo.On(
			"InsertBook", mock.Anything, mock.AnythingOfType("*model.Book"),
		).Return(nil)

		err := mockedRepo.InsertBook(nil, &book)

		assert.Nil(t, err)
		mockedRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockedRepo := createMockedRepository()

		mockedRepo.On(
			"InsertBook", mock.Anything, mock.AnythingOfType("*model.Book"),
		).Return(errors.New("error"))

		err := mockedRepo.InsertBook(nil, &book)

		assert.NotNil(t, err)
	})
}

func createMockedRepository() *mocks.BooksRepo {
	return new(mocks.BooksRepo)
}
