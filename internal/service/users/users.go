package users

import (
	"awesomeProject/internal/domain"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/repository/users/converter"
	"awesomeProject/internal/repository/users/model"
	utils "awesomeProject/pkg/passwords"
	"context"
	"errors"
)

//go:generate mockery --name UsersService
type UsersService struct {
	repo repository.UsersRepo
}

func NewUsersService(repo repository.UsersRepo) *UsersService {
	return &UsersService{repo: repo}
}

func (service *UsersService) SignIn(ctx context.Context, credentials *domain.UserCredentials) (string, error) {

	storedCredentials, err := service.repo.GetUserCredentials(ctx, credentials.Login)

	if err != nil {
		return "", err
	}

	err = utils.ComparePasswords(credentials.Password, storedCredentials.Salt, storedCredentials.Password)

	if err != nil {
		return "", err
	}

	token, err := service.repo.GetJWTToken(ctx, storedCredentials.UserId)

	if err != nil && !errors.Is(err, domain.ErrTokenExpired) {
		return "", err
	}

	if token != "" {
		return token, nil
	}

	token, _ = utils.CreateJwtToken(storedCredentials)
	err = service.repo.SaveJWTToken(ctx, storedCredentials.UserId, token)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *UsersService) SignUp(ctx context.Context, userInfo *domain.SignUpUser) error {
	var emptyCreds model.UserCredentials
	storedCreds, err := service.repo.GetUserCredentials(ctx, userInfo.Login)

	if storedCreds != emptyCreds {
		return domain.ErrLoginAlreadyExists
	}

	userInfo.Salt, err = utils.CreateSalt()
	userInfo.Password = utils.SaltPassword(userInfo.Password, userInfo.Salt)

	if err != nil {
		return err
	}

	userId, err := service.repo.SaveUserCredentials(ctx, converter.ToSignUpInfoFromService(userInfo))

	if err != nil {
		return err
	}

	token, err := utils.CreateJwtToken(storedCreds)

	if err != nil {
		return err
	}

	err = service.repo.SaveJWTToken(ctx, userId, token)

	if err != nil {
		return err
	}

	return nil
}