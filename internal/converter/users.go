package converter

import (
	"awesomeProject/internal/domain"
	desc "awesomeProject/pkg/v1/user"
)

func ToUserCredentialsFromDesc(user *desc.UserCredentials) *domain.UserCredentials {
	return &domain.UserCredentials{
		Login:    user.Login,
		Password: user.Password,
	}
}

func ToSignUpInfoFromDesc(user *desc.SignUpUser) *domain.SignUpUser {
	return &domain.SignUpUser{
		Login:      user.Login,
		Password:   user.Password,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Patronymic: user.Patronymic,
		BirthDate:  user.BirthDate,
	}
}
