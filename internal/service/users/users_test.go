package users

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/service/mocks"
	"testing"
)

func TestSignIn(t *testing.T) {
	credentials := domain.UserCredentials{
		Login:    "admin",
		Password: "admin",
	}

	jwtToken := "jwtToken"

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On(
			"SignIn", mock.Anything, mock.AnythingOfType("*domain.UserCredentials"),
		).Return(jwtToken, nil)

		token, err := mockedService.SignIn(context.TODO(), &credentials)

		assert.Equal(t, token, jwtToken)
		assert.NoError(t, err)

		mockedService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		jwtToken = ""
		expectedErr := domain.ErrInvalidCredentials
		mockedService := createMockedService()

		mockedService.On("SignIn", mock.Anything, mock.AnythingOfType("*domain.UserCredentials")).Return(jwtToken, expectedErr)

		token, err := mockedService.SignIn(context.TODO(), &credentials)

		assert.Equal(t, token, jwtToken)
		assert.Equal(t, expectedErr, err)

		mockedService.AssertExpectations(t)
	})
}

func TestSignUp(t *testing.T) {
	userInfo := domain.SignUpUser{
		Login:      "admin",
		Password:   "admin",
		FirstName:  "admin",
		LastName:   "admin",
		Patronymic: "admin",
		BirthDate:  "21-04-2024",
	}

	t.Run("success", func(t *testing.T) {
		mockedService := createMockedService()

		mockedService.On("SignUp", mock.Anything, mock.AnythingOfType("*domain.SignUpUser")).Return(nil)

		err := mockedService.SignUp(context.TODO(), &userInfo)

		assert.NoError(t, err)
		mockedService.AssertExpectations(t)
	})
}
func createMockedService() *mocks.UsersService {
	return new(mocks.UsersService)
}
