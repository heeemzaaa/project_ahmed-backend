package auth_handler

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) Confirmation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Error parsing the data of the user !"})
		return
	}

	confirmed, err := h.service.Confirmation(&user)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	accessToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.WriteJsonErrors(w, *err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
		Secure:   false, // true in production (HTTPS)
	})

	utils.WriteDataBack(w, map[string]interface{}{
		"message":   "User registered successfully",
		"confirmed": confirmed,
	})
}
