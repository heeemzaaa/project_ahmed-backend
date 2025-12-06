package admin_handler

import (
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AdminHandler) CheckIfAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "There is no cookie related to this user"})
		return
	}
	claims, err := utils.VerifyJWT(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	isAdmin, errJson := h.service.CheckIfAdmin(claims["user_id"].(string))
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	if isAdmin {
		utils.WriteDataBack(w, map[string]string{
			"role": "admin",
		})
	} else {
		utils.WriteDataBack(w, map[string]string{
			"role": "nope",
		})
	}
}
