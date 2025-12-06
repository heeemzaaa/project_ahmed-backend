package videos

import "database/sql"

type VideosRepository struct {
	db *sql.DB
}

func NewVideosRepository(db *sql.DB) *VideosRepository {
	return &VideosRepository{db: db}
}
