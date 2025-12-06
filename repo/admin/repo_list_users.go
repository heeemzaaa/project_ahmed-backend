package admin_repo

import (
	"log"

	"backend/models"
)

func (r *AdminRepository) GetAllUsers() ([]models.User, *models.ErrorJson) {
	rows, err := r.db.Query(`
                SELECT id, first_name, last_name, email, centre, filiere, year,
                       access_premiere_annees, access_deuxieme_annees,
                       access_concours_francais, access_concours_maroc, isAdmin
                FROM user
                ORDER BY last_name, first_name
        `)
	if err != nil {
		log.Println("Error getting all the users: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Centre,
			&user.Filiere,
			&user.Year,
			&user.AccessPremiereAnnees,
			&user.AccessDeuxiemeAnnees,
			&user.AccessConcoursFrancais,
			&user.AccessConcoursMaroc,
			&user.IsAdmin,
		); err != nil {
			log.Println("Error scanning all the users: ", err)
			return nil, &models.ErrorJson{Status: 500, Error:"Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, &models.ErrorJson{Status: 500, Error:"Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	return users, nil
}
