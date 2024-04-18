package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"math/rand"
	"net/http"
	"simple-golang-api/internal/domain"
	desc "simple-golang-api/pkg/v1/order"
	responseFormer "simple-golang-api/pkg/validator"
	"time"
)

func (h *Handler) initOrderRoutes(r chi.Router) {
	r.Route("/order", func(r chi.Router) {
		r.Post("/", h.createOrder())
	})
}

// imitation of an external service
func (h *Handler) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var order desc.Order

		err := json.NewDecoder(r.Body).Decode(&order)

		if err != nil {
			responseFormer.FormErrorResponse(w, r, http.StatusBadRequest, err, domain.ErrInvalidRequest)

			return
		}

		err = responseFormer.IsRequestValid(order)

		if err != nil {
			responseFormer.FormValidationErrorResponse(w, r, http.StatusBadRequest, err, order)

			return
		}

		// create new generator
		randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

		//generate random number in range [1.0, 10.0]
		randomSleep := randSeed.Float64()*9 + 1

		sleepDuration := time.Duration(randomSleep * float64(time.Second))

		time.Sleep(sleepDuration)

		render.JSON(w, r, responseFormer.OK())
	}
}
