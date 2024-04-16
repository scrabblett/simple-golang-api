package users

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/repository/mocks"
	"simple-golang-api/internal/repository/users/model"
	"testing"
)

func TestGetUserCredentials(t *testing.T) {
	login := "admin"

	expectedCredentials := model.UserCredentials{
		Login:    "admin",
		Password: "admin",
		Salt:     "abcde",
		UserId:   1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := createMockedRepository()
		mockUserRepo.On("GetUserCredentials", mock.Anything, login).Return(expectedCredentials, nil)

		creds, err := mockUserRepo.GetUserCredentials(context.Background(), login)

		assert.Equal(t, err, nil)
		assert.Equal(t, creds, expectedCredentials)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("no creds found", func(t *testing.T) {
		expectedErr := domain.ErrInvalidCredentials
		mockUserRepo := createMockedRepository()
		mockUserRepo.On("GetUserCredentials", mock.Anything, login).Return(model.UserCredentials{}, expectedErr)

		_, err := mockUserRepo.GetUserCredentials(context.Background(), login)

		assert.Equal(t, err, expectedErr)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestSaveUserCredentials(t *testing.T) {
	userInfo := model.SignUpUser{
		Login:      "admin",
		Password:   "admin",
		FirstName:  "admin",
		LastName:   "admin",
		Patronymic: "admin",
		BirthDate:  "21-04-2024",
	}

	expectedId := int64(1)

	t.Run("success", func(t *testing.T) {
		mockUserRepo := createMockedRepository()
		mockUserRepo.On("SaveUserCredentials", mock.Anything, &userInfo).Return(expectedId, nil)

		id, err := mockUserRepo.SaveUserCredentials(nil, &userInfo)

		assert.Equal(t, err, nil)
		assert.Equal(t, id, expectedId)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := errors.New("user credentials not inserted")
		mockUserRepo := createMockedRepository()
		mockUserRepo.On(
			"SaveUserCredentials", mock.Anything, &userInfo,
		).Return(int64(0), expectedErr)

		_, err := mockUserRepo.SaveUserCredentials(nil, &userInfo)

		assert.Equal(t, err, expectedErr)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestSaveJWTToken(t *testing.T) {
	userId := int64(1)
	jwt := "jwtToken"

	t.Run("success", func(t *testing.T) {
		mockUserRepo := createMockedRepository()
		mockUserRepo.On(
			"SaveJWTToken", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string"),
		).Return(nil)

		err := mockUserRepo.SaveJWTToken(nil, userId, jwt)

		assert.Equal(t, err, nil)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.UsersRepo)
		expectedErr := errors.New("user credentials not inserted")
		mockUserRepo.On(
			"SaveJWTToken", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("string"),
		).Return(expectedErr)

		err := mockUserRepo.SaveJWTToken(nil, userId, jwt)

		assert.Equal(t, err, expectedErr)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetJWTToken(t *testing.T) {
	expectedToken := "jwtToken"
	userId := int64(1)

	t.Run("success", func(t *testing.T) {
		mockUserRepo := createMockedRepository()

		mockUserRepo.On(
			"GetJWTToken", mock.Anything, mock.AnythingOfType("int64"),
		).Return(expectedToken, nil)

		token, err := mockUserRepo.GetJWTToken(nil, userId)

		assert.Equal(t, token, expectedToken)
		assert.Nil(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		expectedErr := domain.ErrTokenExpired
		expectedToken = ""

		mockUserRepo := createMockedRepository()
		mockUserRepo.On(
			"GetJWTToken", mock.Anything, mock.AnythingOfType("int64"),
		).Return(expectedToken, expectedErr)

		token, err := mockUserRepo.GetJWTToken(nil, userId)

		assert.Equal(t, token, expectedToken)
		assert.Equal(t, err, expectedErr)

		mockUserRepo.AssertExpectations(t)
	})
}

func createMockedRepository() *mocks.UsersRepo {
	return new(mocks.UsersRepo)
}
