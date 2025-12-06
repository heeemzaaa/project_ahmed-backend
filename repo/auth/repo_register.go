package auth_repo

import (
	"database/sql"
	"log"

	"backend/models"
)

// SaveUser inserts a new user record into the database.
func (r *AuthRepository) SaveUser(user *models.User, password string) *models.ErrorJson {
	_, err := r.db.Exec(`
		INSERT INTO user (id, first_name, last_name, email, centre, filiere, password, created_at, is_google, confirmation_code)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, user.ID, user.FirstName, user.LastName, user.Email, user.Centre, user.Filiere, password, user.CreatedAt, user.IsGoogle, user.ConfirmationCode)
	if err != nil {
		log.Println("Error saving the user into the database: ", err)
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}

func (r *AuthRepository) SaveConfirmationCode(user *models.User) error {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM user WHERE id = ? AND confirmation_code IS NOT NULL)",
		user.ID,
	).Scan(&exists)
	if err != nil {
		log.Println("Error checking if the confirmation code is in the database: ", err)
		return err
	}

	if exists {
		_, err := r.db.Exec("UPDATE user SET confirmation_code = ? WHERE id = ?", user.ConfirmationCode, user.ID)
		if err != nil {
			log.Println("Error updating the confirmation code of the user: ", err)
			return err
		}
	}
	return nil
}

// FindByEmail returns a user by email (used for checking duplicates or verifying code).
func (r *AuthRepository) FindByEmail(email string) (bool, *models.ErrorJson) {
	var exists bool
	err := r.db.QueryRow(`SELECT EXISTS (SELECT 1 FROM user WHERE email = ? LIMIT 1)`, email).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error finding the email the user entered: ", err)
			return false, &models.ErrorJson{Status: 401, Error: "L'email que vous avez saisi n'est pas correct, veuillez vérifier à nouveau !"}
		}
		log.Println("Error finding the email the user entered: ", err)
		return false, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	return exists, nil
}
