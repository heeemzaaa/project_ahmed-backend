package auth_service

import (

	"backend/models"
)

func (s *AuthService) Confirmation(user *models.User) (bool, *models.ErrorJson) {
	confirmed, err := s.repo.CheckConfirmation(user)
	if err != nil {
		return false, err
	}

	err = s.repo.SaveDevice(user.ID, user.DeviceId, user.DeviceType)
	if err != nil {
		return false, err
	}

	if confirmed {
		userConfirmed, err := s.repo.SelectUser(user.Email)
		if err != nil {
			return false, err
		}

		err = s.repo.CreateNewSession(userConfirmed.ID, user.DeviceId, user.DeviceType)
		if err != nil {
			return false, err
		}
	}

	return confirmed, nil
}
