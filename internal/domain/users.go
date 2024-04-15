package domain

type UserCredentials struct {
	Login    string
	Password string
}

type SignUpUser struct {
	Login      string
	Password   string
	LastName   string
	FirstName  string
	Patronymic string
	BirthDate  string
	Salt       string
}
