package auth_handler

import (
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) IsConfirmed(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

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

	isConfirmed, errJson := h.service.IsConfirmed(userID)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	utils.WriteDataBack(w, isConfirmed)
}
