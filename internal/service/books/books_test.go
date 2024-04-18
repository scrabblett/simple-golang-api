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

func TestGetBookById(t *testing.T) {
	book := &domain.Book{
		Id:             1,
		Title:          "test",
		Description:    "test",
		AgeGroup:       6,
		PublishingDate: "2005-01-01",
	}

	bookId := book.Id

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"GetBookById", mock.Anything, mock.AnythingOfType("int64"),
		).Return(book, nil)

		expectedBook, err := mockedService.GetBookById(context.TODO(), bookId)

		assert.NoError(t, err)
		assert.Equal(t, expectedBook, book)
	})

	t.Run("failure", func(t *testing.T) {
		emptyBook := &domain.Book{}
		mockedService := createMockedService()

		mockedService.On(
			"GetBookById", mock.Anything, mock.AnythingOfType("int64"),
		).Return(emptyBook, domain.BookNotFound)

		expectedBook, err := mockedService.GetBookById(context.TODO(), 0)

		assert.ErrorIs(t, err, domain.BookNotFound)
		assert.Equal(t, expectedBook, emptyBook)
	})
}

func TestUpdateBookById(t *testing.T) {
	book := &domain.Book{
		Id:             1,
		Title:          "test",
		Description:    "test",
		AgeGroup:       6,
		PublishingDate: "2005-01-01",
	}

	bookId := book.Id

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"UpdateBookById",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("*domain.Book"),
		).Return(nil)

		err := mockedService.UpdateBookById(context.TODO(), bookId, book)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"UpdateBookById",
			mock.Anything,
			mock.AnythingOfType("int64"),
			mock.AnythingOfType("*domain.Book"),
		).Return(domain.BookNotFound)

		err := mockedService.UpdateBookById(context.TODO(), bookId, book)

		assert.ErrorIs(t, err, domain.BookNotFound)
	})
}

func TestDeleteBookById(t *testing.T) {
	bookId := int64(1)
	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"DeleteBookById", mock.Anything, mock.AnythingOfType("int64"),
		).Return(nil)

		err := mockedService.DeleteBookById(context.TODO(), bookId)

		assert.NoError(t, err)
	})

	t.Run("failure", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"DeleteBookById", mock.Anything, mock.AnythingOfType("int64"),
		).Return(domain.BookNotFound)

		err := mockedService.DeleteBookById(context.TODO(), bookId)

		assert.ErrorIs(t, err, domain.BookNotFound)
	})
}

func createMockedService() *mocks.BooksService {
	return new(mocks.BooksService)
}
