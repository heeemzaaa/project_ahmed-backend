package admin_service

import "backend/models"

func (s *AdminService) CheckIfAdmin(userId string) (bool, *models.ErrorJson) {
	return s.repo.CheckIfAdmin(userId)
}
