package auth_service

import (
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) LoginUser(email string, password string, deviceId string, deviceType string) (*models.User, *models.ErrorJson) {
	if email == "" || password == "" {
		return nil, &models.ErrorJson{Status: 400, Error: "Pas assez de données !"}
	}
	
	exists, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	
	if !exists {
		return nil, &models.ErrorJson{Status: 401, Error: "Cette adresse e-mail n'existe pas !"}
	}
	
	user, err := s.repo.SelectUser(email)
	if err != nil {
		return nil, err
	}

	devId, errDev := s.repo.FindDeviceId(user.ID, deviceType)
	if errDev != nil {
		return nil, errDev
	}

	if devId == "" {
		err := s.repo.SaveDevice(user.ID, deviceId, deviceType)
		if err != nil {
			return nil, err
		}
		devId = deviceId
	}

	if devId != deviceId {
		return nil, &models.ErrorJson{Status: 401, Error: "Vous ne pouvez pas accéder au site web depuis ce navigateur, veuillez vérifier avec le professeur Ahmed !"}
	}
	
	if user.IsGoogle {
		return nil, &models.ErrorJson{Status: 400, Error: "Vous êtes inscrit avec un compte Google, vous devez donc vous connecter avec votre compte Google !"}
	}
	
	hashedPassword, err := s.repo.SelectPassword(email)
	if err != nil {
		return nil, err
	}
	
	errPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if errPassword != nil {
		return nil, &models.ErrorJson{Status: 401, Error: "Le mot de passe est incorrect, veuillez réessayer !"}
	}
	
	user.IsConfirmed, err = s.repo.IsConfirmed(user.ID)
	if err != nil {
		return nil, err
	}
	
	if !user.IsConfirmed {
		errConf := s.SendConfirmationEmail(user)
		if errConf != nil {
			return nil, &models.ErrorJson{Status: 500, Error: "Erreur lors de l'envoi d'un nouveau message de confirmation à l'utilisateur !"}
		}
	}
	
	if deviceId == "" || deviceType == "" {
		return nil, &models.ErrorJson{Status: 401, Error: "Vous ne pouvez pas accéder au site web depuis ce navigateur, veuillez vérifier avec le professeur Ahmed !"}
	}
	// here I need to check if the user have the right to enter with the device he have
	err = s.handleUserSessionAccess(user, deviceId, deviceType)
	if err != nil {
		return nil, err
	}

	return user, nil
}
