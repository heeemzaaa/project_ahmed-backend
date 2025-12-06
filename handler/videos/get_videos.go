package videos

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *videosHandler) GetVideosByCategory(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Vous n'avez pas choisi une categorie !"})
		return
	}

	response, err := h.service.GetVideosByCategory(category)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Erreur !!"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	utils.WriteDataBack(w, response)
}

func (h *videosHandler) GetVideoByID(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("id")
	if videoID == "" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Vous n'avez pas choisi un valid video id !"})
		return
	}

	video, err := h.service.GetVideoByID(videoID)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 500, Error: "Erreur !!"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(video)
}
