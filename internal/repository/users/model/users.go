package model

type SignUpUser struct {
	Login      string
	Password   string
	LastName   string
	FirstName  string
	Patronymic string
	BirthDate  string
	Salt       string
}

type UserCredentials struct {
	Login    string
	Password string
	Salt     string
	UserId   int64
}
