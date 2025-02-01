package model

import "time"

type Video struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	VideoID     string    `gorm:"type:VARCHAR(255);uniqueIndex;not null" json:"video_id"` 
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	URL         string    `json:"url"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt time.Time `gorm:"not null;index" json:"published_at"`
}
