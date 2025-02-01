package service

import (
	"context"
	"fmt"
	"go-tube/model"
	"go-tube/storage"
	"log"
	"time"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct {
	DB            storage.DB
	Cache         storage.Cache
	Query         string
	APIKeys       []string 
	round         int // lets keep track of the current API key index and use round robin technique to switch between keys
	errorCount    int 
	maxErrorCount int
}

func NewYoutubeService(db storage.DB, cache storage.Cache, query string, apiKeys []string) *YoutubeService {
	return &YoutubeService{
		DB:            db,
		Cache:         cache,
		Query:         query,
		APIKeys:       apiKeys,
		round:         0,
		errorCount:    0,
		maxErrorCount: 3,
	}
}

func (y *YoutubeService) Start(ctx context.Context, interval time.Duration , cancel context.CancelFunc) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Println("Starting youtube service")
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping youtube service")
			return
		case <-ticker.C:
			if err := y.searchAndStoreVideos(ctx); err != nil {
				log.Printf("Error during video search and store: %v\n", err)
				y.errorCount++
				if y.errorCount >= y.maxErrorCount {
					log.Println("Max error count reached. Stopping youtube service")
					cancel()
					return
				}
			} else {
				y.errorCount = 0
			}
		}
	}
}

func (y *YoutubeService) searchAndStoreVideos(ctx context.Context) error {
	apiKey := y.getNextAPIKey()
	if apiKey == "" {
		return fmt.Errorf("no API keys available")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Println("Failed to create YouTube service:", err)
		return fmt.Errorf("failed to create YouTube service: %v", err)
	}

    publishedAfter := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)
	call := service.Search.List([]string{"id", "snippet"}).
		Q(y.Query).
		Type("video").
		Order("date").
		PublishedAfter(publishedAfter).
		MaxResults(50)
	response, err := call.Do()
	if err != nil {
		log.Printf("Failed to search videos: %v", err)
		return fmt.Errorf("failed to search videos: %v", err)
	}

	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Printf("Failed to parse PublishedAt for video %s: %v\n", item.Id.VideoId, err)
			continue
		}

		video := model.Video{
			VideoID:     item.Id.VideoId,
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			Thumbnail:   item.Snippet.Thumbnails.Default.Url,
			PublishedAt: publishedAt,
		}


		if err := y.DB.InsertVideo(ctx, video); err != nil {
			log.Printf("Failed to insert video %s: %v\n", video.VideoID, err)
			continue
		}
	}

	log.Printf("Successfully fetched and stored %d videos\n", len(response.Items))
	return nil
}

func (y *YoutubeService) getNextAPIKey() string {
	if len(y.APIKeys) == 0 {
		return ""
	}

	apiKey := y.APIKeys[y.round]
	y.round = (y.round + 1) % len(y.APIKeys) 
	return apiKey
}
