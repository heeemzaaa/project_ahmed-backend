package admin_repo

import (
	"log"

	"backend/models"
)

func (r *AdminRepository) CheckIfAdmin(userId string) (bool, *models.ErrorJson) {
	var isAdmin bool
	err := r.db.QueryRow("SELECT isAdmin FROM user WHERE id = ?", userId).Scan(&isAdmin)
	if err != nil {
		log.Println("Error checking if the user is an admin: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Nous ne pouvons pas vérifier si vous êtes un utilisateur, veuillez réessayer plus tard !"}
	}
	return isAdmin, nil
}
