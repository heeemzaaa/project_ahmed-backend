package videos

import service "backend/service/videos"

type videosHandler struct {
	service *service.VideosService
}


func NewVideosHandler(service *service.VideosService) *videosHandler {
	return &videosHandler{service: service}
}

