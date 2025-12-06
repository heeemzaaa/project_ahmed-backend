package auth_service


import repo "backend/repo/auth"

type AuthService struct {
	repo *repo.AuthRepository
}

// NewPostService creates a new service
func NewAuthService(repo *repo.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}