package auth_handler

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) Verify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	cookie, err := r.Cookie("access_token")
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized !"})
		return
	}

	claims, err := utils.VerifyJWT(cookie.Value)
	if err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 401, Error: "Unauthorized because the value doesn't match the one we have!"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid": true,
		"user": map[string]string{
			"id":    claims["user_id"].(string),
			"email": claims["email"].(string),
		},
	})
}
