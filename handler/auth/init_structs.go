package auth_handler

import s "backend/service/auth"

type AuthHandler struct {
	service *s.AuthService
}


func NewAuthHandler(service *s.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}