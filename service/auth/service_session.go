package auth_service

import (
	"backend/models"
)

func (s *AuthService) handleUserSessionAccess(user *models.User, deviceId, deviceType string) *models.ErrorJson {
	session, err := s.repo.GetUserSession(user.ID, deviceType)
	if err != nil {
		return err
	}

	if session != nil {
		// User already has a session for this device type
		if session.DeviceId != deviceId {
			return &models.ErrorJson{Status: 403, Error: "accès refusé : vous ne pouvez en utiliser qu'un seul " + deviceType + " appareil"}
		}

		// Update last_used_at
		err = s.repo.UpdateSessionLastUsed(session.ID)
		if err != nil {
			return err
		}

		return nil
	}

	// No session found — create a new one
	err = s.repo.CreateNewSession(user.ID, deviceId, deviceType)
	if err != nil {
		return err
	}

	return nil
}
