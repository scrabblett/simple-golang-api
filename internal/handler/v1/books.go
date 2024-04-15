package v1

import (
	"awesomeProject/internal/converter"
	"awesomeProject/internal/domain"
	desc "awesomeProject/pkg/v1/book"
	responseFormer "awesomeProject/pkg/validator"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) initBooksRoutes(r chi.Router) {
	r.Route("/books", func(r chi.Router) {
		r.Post("/", h.createBook())
	})
}

func (h *Handler) createBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var book desc.Book

		body := json.NewDecoder(r.Body)

		if err := body.Decode(&book); err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err := responseFormer.IsRequestValid(book)

		if err != nil {
			responseFormer.FormValidationErrorResponse(w, r, http.StatusBadRequest, err, book)

			return
		}

		serviceBook := converter.ToBookFromDesc(&book)

		err = h.services.Books.CreateNewBook(r.Context(), serviceBook)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)

			return
		}

		book.Id = serviceBook.Id

		render.JSON(w, r, book)
	}
}
