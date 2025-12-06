package auth_handler

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid data !"})
		return
	}
	user, err := h.service.RegisterUser(&input)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	utils.WriteDataBack(w, map[string]interface{}{
		"message":           "User registered successfully",
		"user":              user,
		"need_confirmation": !user.IsGoogle,
	})
}
