package auth_repo

import (
	"log"

	"backend/models"

	"github.com/google/uuid"
)

func (r *AuthRepository) SaveDevice(userID string, deviceID string, deviceType string) *models.ErrorJson {
	_, err := r.db.Exec("INSERT INTO user_devices (id, device_id, device_type, user_id) VALUES (?, ?, ?, ?)", uuid.New().String(), deviceID, deviceType, userID)
	if err != nil {
		log.Println("Error saving the device data: ", err)
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}

func (r *AuthRepository) FindDeviceId(userID string, deviceType string) (string, *models.ErrorJson) {
	var deviceID string
	var isThere bool

	err := r.db.QueryRow("SELECT EXISTS (SELECT 1 FROM user_devices WHERE user_id = ? AND device_type = ?)", userID, deviceType).Scan(&isThere)
	if err != nil {
		log.Println("Error checking if the user have a device type: ", err)
		return "", &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	if isThere {
		err = r.db.QueryRow("SELECT device_id FROM user_devices WHERE user_id = ? AND device_type = ?", userID, deviceType).Scan(&deviceID)
		if err != nil {
			log.Println("Error while checking the device type: ", err)
			return "", &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
		}
	}

	return deviceID, nil
}
