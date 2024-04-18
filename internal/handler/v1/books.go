package v1

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"simple-golang-api/internal/converter"
	"simple-golang-api/internal/domain"
	desc "simple-golang-api/pkg/v1/book"
	responseFormer "simple-golang-api/pkg/validator"
	"strconv"
)

func (h *Handler) initBooksRoutes(r chi.Router) {
	r.Route("/book", func(r chi.Router) {
		r.Post("/", h.createBook())
		r.Get("/{id}", h.getBookById())
		r.Put("/{id}", h.updateBook())
		r.Delete("/{id}", h.deleteBook())
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

func (h *Handler) getBookById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bookId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		serviceBook, err := h.services.Books.GetBookById(r.Context(), bookId)

		if err != nil {
			if errors.Is(err, domain.BookNotFound) {
				responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.BookNotFound)

				return
			}

			responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)

			return
		}

		render.JSON(w, r, converter.ToBookFromService(serviceBook))
	}
}

func (h *Handler) updateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bookId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		var book desc.Book

		err = json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err = responseFormer.IsRequestValid(book)

		if err != nil {
			responseFormer.FormValidationErrorResponse(w, r, http.StatusBadRequest, err, book)

			return
		}

		book.Id = bookId
		err = h.services.Books.UpdateBookById(r.Context(), bookId, converter.ToBookFromDesc(&book))

		if err != nil {
			if errors.Is(err, domain.BookNotFound) {
				responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.BookNotFound)

				return
			}
			responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)

			return
		}

		render.JSON(w, r, responseFormer.OK())
	}
}

func (h *Handler) deleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bookId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err = h.services.Books.DeleteBookById(r.Context(), bookId)

		if err != nil {
			if errors.Is(err, domain.BookNotFound) {
				responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.BookNotFound)

				return
			}

			responseFormer.FormErrorResponse(w, r, http.StatusInternalServerError, err, domain.ErrInternalServer)

			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
