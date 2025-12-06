package auth_handler

import (
	"encoding/json"
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 405, Error: "Method not allowed !"})
		return
	}

	type Login struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		DeviceId   string `json:"device_id"`
		DeviceType string `json:"device_type"`
	}

	var login Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		utils.WriteJsonErrors(w, models.ErrorJson{Status: 400, Error: "Error parsing the data of the user !"})
		return
	}

	user, err := h.service.LoginUser(login.Email, login.Password, login.DeviceId, login.DeviceType)
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
		Secure:   false, // set to true in production with HTTPS
	})

	utils.WriteDataBack(w, map[string]interface{}{
		"message": "User logged in successfully",
		"user":    user,
	})
}
