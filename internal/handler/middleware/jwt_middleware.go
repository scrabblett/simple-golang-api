package middleware

import (
	"github.com/go-chi/render"
	"net/http"
	utils "simple-golang-api/pkg/passwords"
	responseFormer "simple-golang-api/pkg/validator"
)

func JwtMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			jwtToken := r.Header.Get("Authorization")

			if jwtToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, responseFormer.Unauthorized())

				return
			}

			ok, err := utils.CheckJwtToken(jwtToken)

			if err != nil || !ok {
				w.WriteHeader(http.StatusUnauthorized)
				render.JSON(w, r, responseFormer.Unauthorized())

				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
