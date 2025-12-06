package auth_service

import "backend/models"

func (s *AuthService) Logout(userId string) *models.ErrorJson {
	err := s.repo.DeleteAllSessions(userId)
	if err != nil {
		return err
	}
	return nil
}
