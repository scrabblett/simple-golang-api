package converter

import (
	"simple-golang-api/internal/domain"
	"simple-golang-api/internal/repository/users/model"
)

func ToSignUpInfoFromService(info *domain.SignUpUser) *model.SignUpUser {
	return &model.SignUpUser{
		Login:      info.Login,
		Password:   info.Password,
		Salt:       info.Salt,
		FirstName:  info.FirstName,
		LastName:   info.LastName,
		Patronymic: info.Patronymic,
		BirthDate:  info.BirthDate,
	}
}

func ToCredentialsFromService(credentials *domain.UserCredentials) *model.UserCredentials {
	return &model.UserCredentials{
		Login:    credentials.Login,
		Password: credentials.Password,
	}
}

func ToCredentialsFromSignUpInfo(info *domain.SignUpUser) *domain.UserCredentials {
	return &domain.UserCredentials{
		Login:    info.Login,
		Password: info.Password,
	}
}
