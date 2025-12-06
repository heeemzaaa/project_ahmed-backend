package videos

import (
	"database/sql"
	"errors"
	"strings"

	"backend/models"
)

// GetByID reads the video row from the database.
// It only selects columns that exist in your videos table; optional metadata is left empty.
func (r *VideosRepository) GetByID(id string) (*models.Video, error) {
	const q = `
	SELECT id, title, category, vdocipherVideoId, orderIndex, folderOrder
	FROM videos
	WHERE id = ?
	LIMIT 1
	`

	row := r.db.QueryRow(q, id)
	var v models.Video
	var folderOrder sql.NullInt64

	err := row.Scan(&v.ID, &v.Title, &v.Category, &v.VDOCipherVideoID, &v.OrderIndex, &folderOrder)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // caller can treat nil as not found
		}
		return nil, err
	}

	v.Title = strings.Split(v.Title, ":::")[1]

	if folderOrder.Valid {
		fo := int(folderOrder.Int64)
		v.FolderOrder = &fo
	} else {
		v.FolderOrder = nil
	}

	// optional fields (Description, Duration, Instructor, Views) are not in DB yet,
	// keep them empty/default. You can extend the query if you add those columns later.

	return &v, nil
}
