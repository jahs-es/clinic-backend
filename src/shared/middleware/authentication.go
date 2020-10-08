package middleware

import (
	"github.com/jahs/clinic-backend/src/shared/auth"
	"net/http"
)

func Authentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := auth.TokenValid(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}
		next(w, r)
}
