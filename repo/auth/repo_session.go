package auth_repo

import (
	"database/sql"
	"log"

	"backend/models"
	"backend/utils"

	"github.com/google/uuid"
)

func (r *AuthRepository) GetUserSession(userID string, deviceType string) (*models.Session, *models.ErrorJson) {
	var session models.Session
	err := r.db.QueryRow(`
		SELECT id, user_id, device_id, device_type, refresh_token, last_used_at, created_at
		FROM user_sessions 
		WHERE user_id = ? AND device_type = ?
	`, userID, deviceType).Scan(
		&session.ID,
		&session.UserId,
		&session.DeviceId,
		&session.DeviceType,
		&session.RefreshToken,
		&session.LastUsedAt,
		&session.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error finding the user session: ", err)
		return nil, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return &session, nil
}

func (r *AuthRepository) CreateNewSession(userID string, deviceId, deviceType string) *models.ErrorJson {
	sessionId := uuid.New().String()

	_, err := r.db.Exec(`
		INSERT INTO user_sessions (id, user_id, device_id, device_type, refresh_token)
		VALUES (?, ?, ?, ?, ?)
	`, sessionId, userID, deviceId, deviceType, utils.GenerateRefreshToken())
	if err != nil {
		log.Println("Error creating a new session for this user: ", err)
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	return nil
}

func (r *AuthRepository) UpdateSessionLastUsed(sessionID string) *models.ErrorJson {
	_, err := r.db.Exec(`
		UPDATE user_sessions 
		SET last_used_at = CURRENT_TIMESTAMP 
		WHERE id = ?
	`, sessionID)
	if err != nil {
		log.Println("Error updating the session: ", err)
		return &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}
	return nil
}
