package routes

import (
	"database/sql"
	"net/http"

	handlerVideos "backend/handler/videos"
	repoVideos "backend/repo/videos"
	serviceVideos "backend/service/videos"
	"backend/middlewares"
)

func VideosRoutes(mux *http.ServeMux, db *sql.DB) *http.ServeMux {
	repo := repoVideos.NewVideosRepository(db)
	service := serviceVideos.NewVideosService(repo)
	handler := handlerVideos.NewVideosHandler(service)

	// Protect both routes using JWT middleware
	mux.Handle("/api/videos", middlewares.NewAuthMiddleware(
		http.HandlerFunc(handler.GetVideosByCategory),
	))

	mux.Handle("/api/video/", middlewares.NewAuthMiddleware(
		http.HandlerFunc(handler.GetVideo),
	))

	return mux
}
