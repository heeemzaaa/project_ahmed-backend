package routes

import (
	handler "backend/handler/auth"
	repo "backend/repo/auth"
	service "backend/service/auth"
	"database/sql"
	"net/http"
)

func AuthRoutes(mux *http.ServeMux, db *sql.DB) (*http.ServeMux, *service.AuthService) {
	authRepo := repo.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/confirmation", authHandler.Confirmation)
	mux.HandleFunc("/api/auth/isconfirmed", authHandler.IsConfirmed)
	mux.HandleFunc("/api/auth/check-user", authHandler.CheckUser)
	mux.HandleFunc("/api/auth/login", authHandler.Login)
	mux.HandleFunc("/api/auth/verify", authHandler.Verify)
	mux.HandleFunc("/api/auth/logout", authHandler.Logout)
	return mux, authService
}


