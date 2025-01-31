package model

import "time"

type Video struct {
	ID          int       `json:"id"`
	VideoID     string    `json:"video_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt time.Time `json:"published_at"`
}