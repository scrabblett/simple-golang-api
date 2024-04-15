package user

type UserCredentials struct {
	Login    string `json:"login" validate:"required,min=5,max=32"`
	Password string `json:"password" validate:"required,min=5,max=32"`
}

type UserInfo struct {
	LastName   string `json:"last_name" validate:"required,min=1,max=32"`
	FirstName  string `json:"first_name" validate:"required,min=1,max=32"`
	Patronymic string `json:"patronymic,omitempty" validate:"min=1,max=32"`
	BirthDate  string `json:"birth_date" validate:"datetime"`
}

type SignUpUser struct {
	Login      string `json:"login" validate:"required,min=5,max=32"`
	Password   string `json:"password" validate:"required,min=5,max=32"`
	LastName   string `json:"last_name" validate:"required,min=1,max=32"`
	FirstName  string `json:"first_name" validate:"required,min=1,max=32"`
	Patronymic string `json:"patronymic,omitempty" validate:"max=32"`
	BirthDate  string `json:"birth_date"`
}
