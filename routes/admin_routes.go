package routes

import (
	"database/sql"
	"net/http"
	"strings"

	handler "backend/handler/admin"
	repo "backend/repo/admin"
	service "backend/service/admin"
	"backend/middlewares"
)

func AdminRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	adminRepo := repo.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	// Protect ALL admin endpoints with JWT middleware
	mux.Handle("/api/admin/users", middlewares.NewAuthMiddleware(
		http.HandlerFunc(adminHandler.ListUsers),
	))

	mux.Handle("/api/auth/me", middlewares.NewAuthMiddleware(
		http.HandlerFunc(adminHandler.CheckIfAdmin),
	))

	mux.Handle("/api/admin/users/", middlewares.NewAuthMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasSuffix(r.URL.Path, "/access") {
				adminHandler.UpdateUserAccess(w, r)
				return
			}
			adminHandler.UpdateUser(w, r)
		}),
	))

	return mux
}
