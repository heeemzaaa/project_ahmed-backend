package auth_repo

import (
	"log"

	"backend/models"
)

func (r *AuthRepository) SelectPassword(email string) (string, *models.ErrorJson) {
	password := ""
	err := r.db.QueryRow("SELECT password FROM user WHERE email = ?", email).Scan(&password)
	if err != nil {
		log.Println("Error in selecting the password: ", err)
		return "", &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return password, nil
}

func (r *AuthRepository) SelectUser(email string) (*models.User, *models.ErrorJson) {
	var user models.User

	err := r.db.QueryRow("SELECT id, first_name, last_name, centre, filiere, is_google FROM user WHERE email = ?", email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Centre,
		&user.Filiere,
		&user.IsGoogle,
	)
	if err != nil {
		log.Println("Error finding the user: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	user.Email = email
	return &user, nil
}
