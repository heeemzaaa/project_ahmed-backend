package admin_handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend/models"
	"backend/utils"
)

func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	userID := strings.TrimPrefix(r.URL.Path, "/api/admin/users/")
	if userID == "" || strings.Contains(userID, "/") {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid user ID"})
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Error parsing the input of the user !"})
		return
	}

	if err := h.service.UpdateUserDetails(userID, &input); err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteDataBack(w, map[string]string{"message": "User updated successfully"})
}
