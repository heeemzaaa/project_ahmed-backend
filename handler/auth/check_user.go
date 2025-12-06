package auth_handler

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) CheckUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Error parsing the email of the user !"})
		return
	}

	exists, err := h.service.CheckUserExists(req.Email)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	resp := struct {
		Exists bool `json:"exists"`
	}{
		Exists: exists,
	}
	w.Header().Set("Content-Type", "application/json")
	utils.WriteDataBack(w, resp)
}
