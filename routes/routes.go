package routes

import (
	"database/sql"
	"net/http"
)

func SetRoutes(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux, _ = AuthRoutes(mux, db)
	mux = AdminRoutes(mux, db)
	mux = VideosRoutes(mux, db)

	return mux
}
