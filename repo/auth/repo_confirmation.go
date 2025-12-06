package auth_repo

import (
	"log"

	"backend/models"
)

func (r *AuthRepository) CheckConfirmation(user *models.User) (bool, *models.ErrorJson) {
	var confirmed string
	var userID string

	err := r.db.QueryRow("SELECT id, confirmation_code FROM user WHERE email = ?", user.Email).Scan(&userID, &confirmed)
	if err != nil {
		log.Println("Error selecting the confirmation code: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}

	}

	if confirmed != user.ConfirmationCode {
		log.Println("Wrong confirmation code: ", err)
		return false, &models.ErrorJson{Status: 401, Error: "Code de confirmation incorrect !"}

	}

	_, err = r.db.Exec("UPDATE user SET is_confirmed = 1 WHERE id = ?", userID)
	if err != nil {
		log.Println("Error updating the confirmation status: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	user.ID = userID
	return true, nil
}

func (r *AuthRepository) IsConfirmed(userID string) (bool, *models.ErrorJson) {
	var confirmed bool
	err := r.db.QueryRow("SELECT is_confirmed FROM user WHERE id = ?", userID).Scan(&confirmed)
	if err != nil {
		log.Println("Error selecting the confirmation code: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}

	}
	return confirmed, nil
}
