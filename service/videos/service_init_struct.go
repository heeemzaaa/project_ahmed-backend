package videos

import (
	"net/http"
	"os"

	"backend/models"
	userrepo "backend/repo/user"
	repo "backend/repo/videos"
)

type VideosService struct {
	repo     *repo.VideosRepository
	userRepo *userrepo.UserRepo
	httpCli  *http.Client
	apiKey   string
}

func NewVideosService(repo *repo.VideosRepository) *VideosService {
	return &VideosService{
		repo:     repo,
		userRepo: userrepo.NewUserRepo(models.DB),
		httpCli:  &http.Client{},
		apiKey:   os.Getenv("VDO_API_SECRET"), // Set this in env securely
	}
}
