package admin_service

import "backend/models"

func (s *AdminService) UpdateUserDetails(userID string, user *models.User) *models.ErrorJson {
	return s.repo.UpdateUserDetails(userID, user)
}
