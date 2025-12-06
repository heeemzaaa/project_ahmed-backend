package auth_handler

import (
	"net/http"

	"backend/models"
	"backend/utils"
)

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
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

	errJson := h.service.Logout(userID)
	if errJson != nil {
		utils.WriteJsonErrors(w, *errJson)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false, // true in production
	})

	utils.WriteDataBack(w, map[string]interface{}{
		"message": "User logged out successfully",
	})
}
