package videos

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"backend/models"
	"backend/utils"
)

func (h *videosHandler) GetVideo(w http.ResponseWriter, r *http.Request) {
	sessionId, err := r.Cookie("access_token")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "There is no access token to this user !"})
		return
	}

	claims, err := utils.VerifyJWT(sessionId.Value)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: err})
		return
	}

	userID := claims["user_id"].(string)

	// extract id from path: expected /api/video/{id}
	path := strings.TrimPrefix(r.URL.Path, "/api/video/")
	if path == "" || path == "/api/video" {
		http.Error(w, "video id required", http.StatusBadRequest)
		return
	}
	videoID := path

	resp, status, err := h.service.GetVideoResponse(context.Background(), userID, videoID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorJson{Status: http.StatusInternalServerError, Error: err.Error()})
		return
	}
	if status == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.ErrorJson{Status: http.StatusNotFound, Error: "video not found"})
		return
	}
	if status == http.StatusForbidden {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(models.ErrorJson{Status: http.StatusForbidden, Error: "access denied"})
		return
	}
	if status == http.StatusUnauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.ErrorJson{Status: http.StatusUnauthorized, Error: "unauthorized"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
