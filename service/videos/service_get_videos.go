package videos

import (
	"backend/models"
)

func (s *VideosService) GetVideosByCategory(category string) (*models.CategoryResponse, error) {
	// Map frontend category IDs to database categories
	var dbCategory string
	var isFoldered bool

	switch category {
	case "1":
		dbCategory = "SUP"
		isFoldered = true
	case "2":
		dbCategory = "SPE"
		isFoldered = true
	case "3":
		dbCategory = "CNC"
		isFoldered = false
	case "4":
		dbCategory = "CF"
		isFoldered = false
	default:
		return nil, nil
	}

	return s.repo.GetVideosByCategory(dbCategory, isFoldered)
}

func (s *VideosService) GetVideoByID(videoID string) (*models.Video, error) {
	return s.repo.GetVideoByID(videoID)
}
