package app

import (
	"banking/errs"
	"banking/service"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	service service.AuthService
}

func (a AuthMiddleware) authorizationHandler() func (handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRouter := mux.CurrentRoute(r)
			currentVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				err := errs.AppError{Code: http.StatusForbidden, Message: "missing token"}
				writeResponse(w, err.Code, err.AsMessage())
				return
			}

			token := getTokenFromHeader(authHeader)
			isAuthorized, err := a.service.Verify(token, currentRouter.GetName(), currentVars)
			if err != nil {
				appErr := errs.AppError{Code: http.StatusForbidden, Message: err.Message}
				writeResponse(w, http.StatusUnauthorized, appErr.AsMessage())
				return
			}

			if isAuthorized {
				next.ServeHTTP(w, r)
			} else {
				err := errs.AppError{Code: http.StatusForbidden, Message: "missing token"}
				writeResponse(w, err.Code, err.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
