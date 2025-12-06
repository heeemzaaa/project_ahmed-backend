package admin_repo

import "backend/models"

func (r *AdminRepository) UpdateUserDetails(userID string, user *models.User) *models.ErrorJson {
	_, err := r.db.Exec(`
                UPDATE user
                SET first_name = ?,
                    last_name = ?,
                    email = ?,
                    centre = ?,
                    filiere = ?,
                    year = ?
                WHERE id = ?
        `, user.FirstName, user.LastName, user.Email, user.Centre, user.Filiere, user.Year, userID)
	if err != nil {
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}
