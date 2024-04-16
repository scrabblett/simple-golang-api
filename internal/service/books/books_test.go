package books

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/service/mocks"
	"testing"
	"time"
)

func TestCreateNewBook(t *testing.T) {
	book := domain.Book{
		Title:          "Test",
		Description:    "Test",
		AgeGroup:       18,
		PublishingDate: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"CreateNewBook", mock.Anything, mock.AnythingOfType("*domain.Book"),
		).Return(nil)

		err := mockedService.CreateNewBook(nil, &book)

		assert.NoError(t, err)
		mockedService.AssertExpectations(t)
	})
}

func createMockedService() *mocks.BooksService {
	return new(mocks.BooksService)
}
