package models

// Video represents a video row from the DB (plus fields used by the API)
type Video struct {
	ID               string `json:"id"`
	Title            string `json:"title"`
	Category         string `json:"category"`
	VDOCipherVideoID string `json:"vdocipherVideoId"`
	OrderIndex       int    `json:"orderIndex"`
	FolderOrder      *int   `json:"folderOrder,omitempty"`
	Description      string `json:"description,omitempty"`
	Duration         string `json:"duration,omitempty"`
	Instructor       string `json:"instructor,omitempty"`
	Views            int    `json:"views,omitempty"`
}

// VideoResponse is the shape returned to the frontend
type VideoResponse struct {
	Title                 string `json:"title"`
	Description           string `json:"description"`
	Duration              string `json:"duration"`
	Instructor            string `json:"instructor"`
	Category              string `json:"category"`
	Views                 int    `json:"views"`
	VdoCipherOTP          string `json:"vdocipher_otp"`
	VdoCipherPlaybackInfo string `json:"vdocipher_playback_info"`
	UserId                string `json:"user_id"`
	FullName              string `json:"full_name"`
}

type Folder struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Icon       string  `json:"icon"`
	Videos     []Video `json:"videos"`
	OrderIndex int     `json:"orderIndex"`
}

type CategoryResponse struct {
	Folders []Folder `json:"folders"`
	Videos  []Video  `json:"videos"`
}
