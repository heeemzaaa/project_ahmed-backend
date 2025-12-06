package admin_service

import repo "backend/repo/admin"

type AdminService struct {
	repo *repo.AdminRepository
}

func NewAdminService(repo *repo.AdminRepository) *AdminService {
	return &AdminService{repo: repo}
}
