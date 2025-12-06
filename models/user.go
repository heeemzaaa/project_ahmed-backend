package models

import "time"

type User struct {
	ID                     string `json:"id"`
	FirstName              string `json:"first_name"`
	LastName               string `json:"last_name"`
	Email                  string `json:"email"`
	Centre                 string `json:"centre"`
	Filiere                string `json:"filiere"`
	Year                   string `json:"year"`
	AccessPremiereAnnees   bool   `json:"access_premiere_annees"`
	AccessDeuxiemeAnnees   bool   `json:"access_deuxieme_annees"`
	AccessConcoursFrancais bool   `json:"access_concours_francais"`
	AccessConcoursMaroc    bool   `json:"access_concours_maroc"`
	Password               string `json:"password"`
	CreatedAt              time.Time
	IsGoogle               bool    `json:"is_google"`
	IsAdmin                bool    `json:"is_admin"`
	DeviceId               string  `json:"device_id"`
	DeviceType             string  `json:"device_type"`
	Session                Session `json:"session"`
	ConfirmationCode       string  `json:"confirmation_code"`
	IsConfirmed            bool    `json:"is_confirmed"`
}

type Session struct {
	ID           string    `json:"id"`
	UserId       string    `json:"user_id"`
	DeviceId     string    `json:"device_id"`
	DeviceType   string    `json:"device_type"`
	RefreshToken string    `json:"refresh_token"`
	LastUsedAt   time.Time `json:"last_used_at"`
	CreatedAt    time.Time `json:"created_at"`
}
