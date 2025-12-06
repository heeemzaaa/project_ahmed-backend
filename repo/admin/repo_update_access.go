package admin_repo

import (
	"log"

	"backend/models"
)

type AccessUpdate struct {
	AccessPremiereAnnees   bool
	AccessDeuxiemeAnnees   bool
	AccessConcoursFrancais bool
	AccessConcoursMaroc    bool
}

func (r *AdminRepository) UpdateUserAccess(userID string, access AccessUpdate) *models.ErrorJson {
	_, err := r.db.Exec(`
                UPDATE user
                SET access_premiere_annees = ?,
                    access_deuxieme_annees = ?,
                    access_concours_francais = ?,
                    access_concours_maroc = ?
                WHERE id = ?
        `, access.AccessPremiereAnnees, access.AccessDeuxiemeAnnees, access.AccessConcoursFrancais, access.AccessConcoursMaroc, userID)
	if err != nil {
		log.Println("Error updating the user access: ", err)
        return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}
