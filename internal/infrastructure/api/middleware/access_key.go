package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ViniciusMartinsS/manager/internal/controller/common"
)

const (
	ROUTE_BY_PASS_MIDDLEWARE = "/auth/login"
	BEARER                   = "Bearer "
)

func CheckAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == ROUTE_BY_PASS_MIDDLEWARE {
			next.ServeHTTP(w, r)
			return
		}

		authorization := r.Header["Authorization"]
		if len(authorization) == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if authorization[0] == "" && strings.Contains(authorization[0], BEARER) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		accessToken := strings.ReplaceAll(authorization[0], BEARER, "")

		isValid, claims := common.IsAccessTokenValid(accessToken)
		if !isValid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		r.Header.Add("User", strconv.Itoa(claims.UserId))
		next.ServeHTTP(w, r)
	})
}
