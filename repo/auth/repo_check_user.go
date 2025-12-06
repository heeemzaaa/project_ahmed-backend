package auth_repo

import (
	"database/sql"
	"log"

	"backend/models"
)

func (r *AuthRepository) UserExistsByEmail(email string) (bool, *models.ErrorJson) {
	var user models.User
	row := r.db.QueryRow("SELECT email FROM user WHERE email = ?", email)
	err := row.Scan(&user.Email)

	if err == sql.ErrNoRows {
		log.Println("Error there is no row compatible to the one you're looking for")
		return false, &models.ErrorJson{Status: 400, Error: "Il n'existe aucun compte avec cet e-mail, veuillez vérifier à nouveau !"}
	}
	if err != nil {
		log.Println("Error checking if the user exist by email: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return true, nil
}
