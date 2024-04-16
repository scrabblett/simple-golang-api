package books

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/service/mocks"
	"testing"
)

func TestCreateNewBook(t *testing.T) {
	book := domain.Book{
		Title:          "Test",
		Description:    "Test",
		AgeGroup:       18,
		PublishingDate: "01-01-2000",
	}

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"CreateNewBook", mock.Anything, mock.AnythingOfType("*domain.Book"),
		).Return(nil)

		err := mockedService.CreateNewBook(context.TODO(), &book)

		assert.NoError(t, err)
		mockedService.AssertExpectations(t)
	})
}

func createMockedService() *mocks.BooksService {
	return new(mocks.BooksService)
}
