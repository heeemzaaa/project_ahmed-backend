package videos

import (
	"fmt"
	"strings"

	"backend/models"
)

func (r *VideosRepository) GetVideosByCategory(dbCategory string, isFoldered bool) (*models.CategoryResponse, error) {
	response := &models.CategoryResponse{
		Folders: []models.Folder{},
		Videos:  []models.Video{},
	}

	if isFoldered {
		// Get distinct folder orders
		folderRows, err := r.db.Query(`
			SELECT DISTINCT folderOrder, title 
			FROM videos 
			WHERE category = ? AND folderOrder IS NOT NULL 
			ORDER BY folderOrder ASC
		`, dbCategory)
		if err != nil {
			return nil, err
		}
		defer folderRows.Close()

		for folderRows.Next() {
			var folderOrder int
			var folderTitle string
			if err := folderRows.Scan(&folderOrder, &folderTitle); err != nil {
				continue
			}
			folderSplitted := strings.Split(folderTitle, ":::")
			folderName := folderSplitted[0]

			folder := models.Folder{
				ID:         fmt.Sprintf("folder_%d", folderOrder),
				Name:       folderName,
				OrderIndex: folderOrder,
			}

			// Assign icon based on category
			if dbCategory == "SUP" {
				folder.Icon = "📚"
			} else {
				folder.Icon = "🎯"
			}

			// Get videos in this folder
			videoRows, err := r.db.Query(`
				SELECT id, title, vdocipherVideoId, orderIndex 
				FROM videos 
				WHERE category = ? AND folderOrder = ? 
				ORDER BY orderIndex ASC
			`, dbCategory, folderOrder)
			if err != nil {
				continue
			}

			for videoRows.Next() {
				var video models.Video
				if err := videoRows.Scan(&video.ID, &video.Title, &video.VDOCipherVideoID, &video.OrderIndex); err != nil {
					continue
				}
				video.Title = strings.Split(video.Title, ":::")[1]
				folder.Videos = append(folder.Videos, video)
			}
			videoRows.Close()
			var exists bool
			for i := 0; i < len(response.Folders); i++ {
				if folder.OrderIndex == response.Folders[i].OrderIndex {
					exists = true
					break
				}
			}

			if !exists {
				response.Folders = append(response.Folders, folder)
			}

		}
	} else {
		// Get videos directly (no folders)
		rows, err := r.db.Query(`
		SELECT id, title, vdocipherVideoId, orderIndex 
		FROM videos 
		WHERE category = ? 
		ORDER BY orderIndex ASC
		`, dbCategory)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var video models.Video
			if err := rows.Scan(&video.ID, &video.Title, &video.VDOCipherVideoID, &video.OrderIndex); err != nil {
				continue
			}
			response.Videos = append(response.Videos, video)
		}
	}

	return response, nil
}

func (r *VideosRepository) GetVideoByID(videoID string) (*models.Video, error) {
	var video models.Video
	err := r.db.QueryRow(`
		SELECT id, title, category, vdocipherVideoId, orderIndex 
		FROM videos 
		WHERE id = ?
	`, videoID).Scan(&video.ID, &video.Title, &video.Category, &video.VDOCipherVideoID, &video.OrderIndex)
	if err != nil {
		return nil, err
	}
	return &video, nil
}
