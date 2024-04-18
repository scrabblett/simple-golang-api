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
