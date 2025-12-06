package auth_service

import (
	"fmt"
	"math/rand"
	"time"

	"backend/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func (s *AuthService) RegisterUser(user *models.User) (*models.User, *models.ErrorJson) {
	existing, _ := s.repo.FindByEmail(user.Email)
	if existing {
		return nil, &models.ErrorJson{Status: 409, Error: "Cet e-mail est déjà utilisé !"}
	}

	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	if !user.IsGoogle {
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
	}

	errJson := s.repo.SaveUser(user, string(hashedPassword))
	if errJson != nil {
		fmt.Println("Error confiramtion 11: ", errJson)
		return nil, errJson
	}

	// Send confirmation email if not Google
	if !user.IsGoogle {
		if err := s.SendConfirmationEmail(user); err != nil {
			fmt.Println("Error confiramtion: ", err)
			return nil, &models.ErrorJson{Status: 500, Error: "Oups, un problème est survenu de notre côté, veuillez réessayer plus tard !"}
		}
	}

	return user, nil
}

// sendConfirmationEmail sends the confirmation code to the user's email.
func (s *AuthService) SendConfirmationEmail(user *models.User) error {
	user.ConfirmationCode = fmt.Sprintf("%06d", rand.Intn(1000000))
	err := s.repo.SaveConfirmationCode(user)
	if err != nil {
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "hamzaelkhawlani00@gmail.com")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Confirm your account")
	m.SetBody("text/plain", fmt.Sprintf(`
Bonjour %s,

Merci pour votre inscription !

Voici votre code de confirmation : %s

Veuillez saisir ce code dans l'application pour activer votre compte.

Cordialement,
L'équipe ProfAhmedCpge
`, user.FirstName, user.ConfirmationCode))

	d := gomail.NewDialer("smtp.gmail.com", 587, "hamzaelkhawlani00@gmail.com", "zfid tjef fnqt fncp")

	return d.DialAndSend(m)
}
