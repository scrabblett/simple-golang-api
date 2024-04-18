package validator

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"reflect"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Errors []ValidationError `json:"errors"`
}

const (
	StatusError = "error"
)

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func OK() Response {
	return Response{
		Status: "ok",
	}
}

func Unauthorized() Response {
	return Response{
		Status: "unauthorized",
	}
}

func FormErrorResponse(w http.ResponseWriter, r *http.Request, sc int, logErr error, outputErr error) {
	w.WriteHeader(sc)
	zap.L().Error(logErr.Error())
	render.JSON(w, r, Error(outputErr.Error()))
}

func FormValidationErrorResponse(w http.ResponseWriter, r *http.Request, sc int, validationError error, st interface{}) {
	var outputErrors ValidationErrorResponse
	var validationErrors validator.ValidationErrors

	if errors.As(validationError, &validationErrors) {
		for _, err := range validationErrors {

			/// find json-name of structure field
			t := reflect.TypeOf(st)
			field, _ := t.FieldByName(err.Field())
			jsonName, _ := field.Tag.Lookup("json")

			/// form json response with validation errors
			ve := ValidationError{
				Field:   jsonName,
				Message: generateValidateMessage(err),
			}

			outputErrors.Errors = append(outputErrors.Errors, ve)
		}

		w.WriteHeader(sc)
		zap.L().Error(validationError.Error())
		render.JSON(w, r, outputErrors)
	}
}

func IsRequestValid(str interface{}) error {
	validate := validator.New()

	err := validate.Struct(str)

	if err != nil {
		return err
	}

	return err
}

var validationErrorMessage = map[string]string{
	"required": "field is required",
	"min":      "value is too short",
	"max":      "value is too long",
}

func generateValidateMessage(error validator.FieldError) string {
	errorMessage := validationErrorMessage[error.Tag()]

	return errorMessage
}
