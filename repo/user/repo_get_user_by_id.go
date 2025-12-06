package userrepo

import (
	"database/sql"
	"errors"

	"backend/models"
)


// GetByID fetches user row and fills access flags
func (r *UserRepo) GetByID(id string) (*models.User, error) {
	const q = `
	SELECT id, first_name, last_name, email,
	       access_premiere_annees, access_deuxieme_annees,
	       access_concours_francais, access_concours_maroc
	FROM user
	WHERE id = ?
	LIMIT 1
	`

	row := r.db.QueryRow(q, id)

	var u models.User
	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.Email,
		&u.AccessPremiereAnnees, &u.AccessDeuxiemeAnnees,
		&u.AccessConcoursFrancais, &u.AccessConcoursMaroc,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
