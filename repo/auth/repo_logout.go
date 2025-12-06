package auth_repo

import (
	"log"
	"backend/models"
)

func (r *AuthRepository) DeleteAllSessions(userId string) *models.ErrorJson {
	_, err := r.db.Exec("DELETE FROM user_sessions WHERE user_id = ?", userId)
	if err != nil {
		log.Println("error deleting the user session from the database: ", err)
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}
