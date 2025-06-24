package handler

import (
	"encoding/json"
	"go-23/dto"
	"go-23/service"
	"go-23/utils"
	"net/http"
)

type AuthHandler struct {
	Service service.Service
}

func NewAuthHandler(service service.Service) AuthHandler {
	return AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var LoginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&LoginRequest); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// validate
	message, err := utils.ValidateData(LoginRequest)
	if err != nil {
		utils.ResponseBadRequest2(w, http.StatusBadRequest, message)
		return
	}

	user, err := h.Service.AuthService.Login(LoginRequest.Email, LoginRequest.Password)
	if err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "Login success", user)
}
