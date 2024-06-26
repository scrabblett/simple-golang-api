//go:build integration_test

package handler

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"net/url"
	"simple-golang-api/internal/repository/users/model"
	utils "simple-golang-api/pkg/passwords"
	desc "simple-golang-api/pkg/v1/book"
	"testing"
)

const (
	booksPath = basePath + "/book"
)

func TestCreateBook(t *testing.T) {
	jwtToken, _ := utils.CreateJwtToken(model.UserCredentials{})
	book := desc.Book{
		Title:          "Test",
		Description:    "Test",
		AgeGroup:       18,
		PublishingDate: "01-01-2000",
	}

	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.New(t, u.String())

	t.Run("success", func(t *testing.T) {
		req := e.POST(booksPath).WithJSON(book).WithHeader("Authorization", jwtToken).Expect().Status(http.StatusOK).JSON().Object()

		schema := `{
		  "$schema": "http://json-schema.org/draft-07/schema#",
		  "title": "Generated schema for Root",
		  "type": "object",
		  "properties": {
			"title": {
			  "type": "string"
			},
			"description": {
			  "type": "string"
			},
			"age_group": {
			  "type": "number"
			},
			"publishing_date": {
			  "type": "string"
			}
		  },
		  "required": [
			"title",
			"description",
			"age_group",
			"publishing_date"
		  ]
		}`

		req.Schema(schema)
	})

	t.Run("unauthorized", func(t *testing.T) {
		req := e.POST(booksPath).WithJSON(book).Expect().Status(http.StatusUnauthorized).JSON().Object()

		req.Value("status").String().Contains("unauthorized")
	})
}
