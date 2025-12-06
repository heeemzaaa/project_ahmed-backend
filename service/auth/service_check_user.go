package auth_service

import "backend/models"

func (s *AuthService) CheckUserExists(email string) (bool, *models.ErrorJson) {
	return s.repo.UserExistsByEmail(email)
}
