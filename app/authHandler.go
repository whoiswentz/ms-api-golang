package app

import (
	"banking/dto"
	"banking/logger"
	"banking/service"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service service.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var lr dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&lr); err != nil {
		logger.Error("error while decoding login request: " + err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Login(lr)
	if err != nil {
		writeResponse(w, http.StatusUnauthorized, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, token)
}
