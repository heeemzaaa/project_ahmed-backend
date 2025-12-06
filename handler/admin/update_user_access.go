package admin_handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend/models"
	service "backend/service/admin"
	"backend/utils"
)

func (h *AdminHandler) UpdateUserAccess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	trimmed := strings.TrimPrefix(r.URL.Path, "/api/admin/users/")
	parts := strings.Split(trimmed, "/")
	if len(parts) != 2 || parts[0] == "" || parts[1] != "access" {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Invalid access path"})
		return
	}

	var input service.AccessUpdateInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Error parsing the data of the user !"})
		return
	}

	if err := h.service.UpdateUserAccess(parts[0], input); err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.WriteDataBack(w, map[string]string{"message": "Access updated successfully"})
}
