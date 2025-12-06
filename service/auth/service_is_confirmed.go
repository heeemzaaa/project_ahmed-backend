package auth_service

import "backend/models"

func (s *AuthService) IsConfirmed(userId string) (bool, *models.ErrorJson) {
	IsConfirmed, err := s.repo.IsConfirmed(userId)
	if err != nil {
		return false, err
	}
	return IsConfirmed, nil
}
