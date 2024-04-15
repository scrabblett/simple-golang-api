package users

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/repository/mocks"
	"awesomeProject/internal/repository/users/model"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetUserCredentials(t *testing.T) {
	mockUserRepo := new(mocks.UsersRepo)

	login := "admin"

	expectedCredentials := model.UserCredentials{
		Login:    "admin",
		Password: "admin",
		Salt:     "abcde",
		UserId:   1,
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetUserCredentials", mock.Anything, login).Return(expectedCredentials, nil)

		mockUserRepo.GetUserCredentials(context.Background(), login)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("no creds found", func(t *testing.T) {
		mockUserRepo.On("GetUserCredentials", mock.Anything, login).Return(model.UserCredentials{}, domain.ErrInvalidCredentials)

		mockUserRepo.GetUserCredentials(context.Background(), login)

		mockUserRepo.AssertExpectations(t)
	})
}

func TestSaveUserCredentials(t *testing.T) {
	mockUserRepo := new(mocks.UsersRepo)

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
		mockUserRepo.On("SaveUserCredentials", mock.Anything, &userInfo).Return(expectedId, nil)

		mockUserRepo.SaveUserCredentials(nil, &userInfo)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo.On("SaveUserCredentials", mock.Anything, &userInfo).Return(0, errors.New("user credentials not inserted"))

		mockUserRepo.SaveUserCredentials(nil, &userInfo)
	})
}
