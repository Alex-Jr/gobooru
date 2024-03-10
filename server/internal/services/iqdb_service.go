package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobooru/internal/models"
	"io"
	"net/http"
	"time"
)

type IQDBService interface {
	HandlePost(post models.Post) ([]models.PostRelation, error)
}

type iqdbService struct {
	IQDB_URL  string
	BASE_PATH string
}

type IQDBServiceConfig struct {
	IQDB_URL  string
	BASE_PATH string
}

func NewIQDBService(c IQDBServiceConfig) IQDBService {
	return &iqdbService{
		IQDB_URL:  c.IQDB_URL,
		BASE_PATH: c.BASE_PATH,
	}
}

func (s *iqdbService) HandlePost(post models.Post) ([]models.PostRelation, error) {
	url := fmt.Sprintf("%s/add", s.IQDB_URL)

	body, err := json.Marshal(map[string]interface{}{
		"postId": post.ID,
		// using thumb path because iqdb only accepts jpg, png, and webp
		"filePath": fmt.Sprintf("%s%s", s.BASE_PATH, post.ThumbPath),
	})

	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Perform the POST request
	client := &http.Client{
		// Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("iqdb returned status code %d", resp.StatusCode)
	}

	type AddResponseBody struct {
		Similarities []struct {
			PostID     int     `json:"postId"`
			Similarity float32 `json:"similarity"`
		} `json:"similarities"`
	}

	var responseBody AddResponseBody

	err = json.Unmarshal(response, &responseBody)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	postRelations := make([]models.PostRelation, 0, len(responseBody.Similarities))

	now := time.Now()

	for _, similarity := range responseBody.Similarities {
		if similarity.Similarity < 60 {
			continue
		}

		postRelations = append(postRelations, models.PostRelation{
			CreatedAt:   now,
			PostID:      post.ID,
			OtherPostID: similarity.PostID,
			Similarity:  int(similarity.Similarity * 100),
			Type:        "SIMILAR",
		})
	}

	return postRelations, nil
}
