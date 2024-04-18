package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	userConverter "simple-golang-api/internal/converter"
	"simple-golang-api/internal/domain"
	desc "simple-golang-api/pkg/v1/user"
	responseFormer "simple-golang-api/pkg/validator"
)

func (h *Handler) initAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", h.authUser())
		r.Post("/register", h.signUp())
	})
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (h *Handler) authUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userCredentials desc.UserCredentials

		body := json.NewDecoder(r.Body)

		if err := body.Decode(&userCredentials); err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err := responseFormer.IsRequestValid(userCredentials)

		if err != nil {
			responseFormer.FormValidationErrorResponse(w, r, http.StatusBadRequest, err, userCredentials)

			return
		}

		convertedInfo := userConverter.ToUserCredentialsFromDesc(&userCredentials)

		token, err := h.services.Users.SignIn(r.Context(), convertedInfo)

		if err != nil {
			if errors.Is(err, domain.ErrInvalidCredentials) {
				responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidCredentials)
			} else {
				responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)
			}

			return
		}

		render.JSON(w, r, AuthResponse{Token: token})
	}
}

func (h *Handler) signUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newUser desc.SignUpUser
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err = responseFormer.IsRequestValid(newUser)

		if err != nil {
			responseFormer.FormValidationErrorResponse(w, r, http.StatusBadRequest, err, newUser)

			return
		}

		convertedBody := userConverter.ToSignUpInfoFromDesc(&newUser)
		err = h.services.Users.SignUp(r.Context(), convertedBody)

		if err != nil {
			if errors.Is(err, domain.ErrLoginAlreadyExists) {
				responseFormer.FormErrorResponse(w, r, http.StatusUnprocessableEntity, err, domain.ErrLoginAlreadyExists)
			} else {
				responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)
			}

			return
		}

		render.JSON(w, r, responseFormer.OK())
	}
}
