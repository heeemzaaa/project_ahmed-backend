package admin_handler

import (
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AdminHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	users, err := h.service.ListUsers()
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		http.Error(w, "Failed to load users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteDataBack(w, users)
}
