//go:build integration_test

package handler

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gavv/httpexpect"
	"net/http"
	"net/url"
	desc "simple-golang-api/pkg/v1/user"
	"testing"
	"time"
)

const (
	host         = "localhost:8080"
	basePath     = "/api/v1"
	authPath     = "/auth"
	loginPath    = basePath + authPath + "/login"
	registerPath = basePath + authPath + "/register"
	timeFormat   = "2006-01-02 15:04:05"
)

func TestAuthUser(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.New(t, u.String())

	t.Run("success", func(t *testing.T) {
		e.POST(loginPath).WithJSON(desc.UserCredentials{
			Login:    "admin",
			Password: "admin",
		}).Expect().Status(http.StatusOK).Body().NotEmpty()
	})

	t.Run("invalid credentials", func(t *testing.T) {
		req := e.POST(loginPath).WithJSON(desc.UserCredentials{
			Login:    "admin",
			Password: "12345",
		}).Expect().Status(http.StatusBadRequest).JSON().Object()

		req.Value("error").String().Contains("invalid credentials")
	})
}

func TestAuthRegister(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.New(t, u.String())

	t.Run("success", func(t *testing.T) {
		e.POST(registerPath).WithJSON(desc.SignUpUser{
			Login:      gofakeit.LoremIpsumSentence(2),
			Password:   gofakeit.LoremIpsumSentence(2),
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Patronymic: gofakeit.LoremIpsumWord(),
			BirthDate:  time.Now().AddDate(-18, 0, 0).Format(timeFormat),
		}).Expect().Status(http.StatusOK).Body().NotEmpty()
	})

	t.Run("user exists", func(t *testing.T) {
		req := e.POST(registerPath).WithJSON(desc.SignUpUser{
			Login:      "admin",
			Password:   "admin",
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Patronymic: gofakeit.LoremIpsumWord(),
			BirthDate:  time.Now().AddDate(-18, 0, 0).Format(timeFormat),
		}).Expect().Status(http.StatusUnprocessableEntity).JSON().Object()

		req.Value("error").String().Contains("login already exists")
	})
}
