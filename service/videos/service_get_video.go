package videos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"backend/models"
)

type vdoOtpResponse struct {
	OTP          string          `json:"otp"`
	PlaybackInfo string `json:"playbackInfo"`
}

// mapping of category to user access column (confirmed by you)
func (s *VideosService) categoryAllowedForUser(v *models.Video, u *models.User) bool {
	switch v.Category {
	case "SUP":
		return u.AccessPremiereAnnees
	case "SPE":
		return u.AccessDeuxiemeAnnees
	case "concours_marocains", "CM", "cnc", "CNC":
		return u.AccessConcoursMaroc
	case "concours_francais", "CF":
		return u.AccessConcoursFrancais
	default:
		// unknown category -> deny by default
		return false
	}
}

// GetVideoResponse checks user access, obtains vdocipher OTP, and returns the payload for frontend.
func (s *VideosService) GetVideoResponse(ctx context.Context, userID string, videoID string) (*models.VideoResponse, int, error) {
	// 1. Get video
	v, err := s.repo.GetByID(videoID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if v == nil {
		return nil, http.StatusNotFound, nil
	}

	// 2. Get user
	u, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if u == nil {
		return nil, http.StatusUnauthorized, nil
	}

	// 3. Verify access
	if !s.categoryAllowedForUser(v, u) {
		return nil, http.StatusForbidden, nil
	}

	// 4. Call VdoCipher to create OTP
	if s.apiKey == "" {
		// If API key is not set, return server error
		return nil, http.StatusInternalServerError, fmt.Errorf("VdoCipher API key not configured (VDO_API_SECRET)")
	}

	otp, playbackInfoStr, err := s.requestVdoOtp(ctx, v.VDOCipherVideoID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// 5. Build response payload
	resp := &models.VideoResponse{
		Title:                 v.Title,
		Description:           v.Description,
		Duration:              v.Duration,
		Instructor:            v.Instructor,
		Category:              v.Category,
		Views:                 v.Views,
		VdoCipherOTP:          otp,
		VdoCipherPlaybackInfo: playbackInfoStr,
		UserId: userID,
		FullName: u.FirstName + " " + u.LastName,
	}

	return resp, http.StatusOK, nil
}

func (s *VideosService) requestVdoOtp(ctx context.Context, vdoID string) (string, string, error) {
	url := fmt.Sprintf("https://dev.vdocipher.com/api/videos/%s/otp", vdoID)
	// payload can be extended if required (e.g. ttl, externalId)
	body := map[string]interface{}{}
	payload, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(payload))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Apisecret "+s.apiKey)

	res, err := s.httpCli.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		bts, _ := ioutil.ReadAll(res.Body)
		return "", "", fmt.Errorf("vdocipher error: status=%d body=%s", res.StatusCode, string(bts))
	}

	var vdoResp vdoOtpResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&vdoResp); err != nil {
		return "", "", err
	}

	playbackInfoStr := vdoResp.PlaybackInfo

	// // playbackInfo might be an object — serialize it to compact string to send to frontend
	// playbackInfoBytes, _ := json.Marshal(vdoResp.PlaybackInfo)
	// playbackInfoStr := string(playbackInfoBytes)
	// if playbackInfoStr == "null" {
	// 	// if playbackInfo is raw null, send empty object
	// 	playbackInfoStr = "{}"
	// }

	return vdoResp.OTP, playbackInfoStr, nil
}
