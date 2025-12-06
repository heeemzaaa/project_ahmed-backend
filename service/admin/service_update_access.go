package admin_service

import (
	"backend/models"
	repo "backend/repo/admin"
)

type AccessUpdateInput struct {
	AccessPremiereAnnees   bool `json:"access_premiere_annees"`
	AccessDeuxiemeAnnees   bool `json:"access_deuxieme_annees"`
	AccessConcoursFrancais bool `json:"access_concours_francais"`
	AccessConcoursMaroc    bool `json:"access_concours_maroc"`
}

func (s *AdminService) UpdateUserAccess(userID string, access AccessUpdateInput) *models.ErrorJson {
	return s.repo.UpdateUserAccess(userID, repo.AccessUpdate{
		AccessPremiereAnnees:   access.AccessPremiereAnnees,
		AccessDeuxiemeAnnees:   access.AccessDeuxiemeAnnees,
		AccessConcoursFrancais: access.AccessConcoursFrancais,
		AccessConcoursMaroc:    access.AccessConcoursMaroc,
	})
}
