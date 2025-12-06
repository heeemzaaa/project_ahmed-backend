package admin_service

import "backend/models"

func (s *AdminService) ListUsers() ([]models.User, *models.ErrorJson) {
	return s.repo.GetAllUsers()
}
