package service

import (
	"context"
	"go-tube/storage"
	"log"
	"time"
)

type YoutubeService struct {
	DB      storage.DB
	Cache   storage.Cache
	Query   string
	APIKeys []string
	round   int
}

func NewYoutubeService(db storage.DB, cache storage.Cache, query string , apiKeys []string) *YoutubeService {
	return &YoutubeService{
		DB:    db,
		Cache: cache,
		Query: query,
		APIKeys: apiKeys,
		round: 0,
	}
}

func (y *YoutubeService) Start(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Println("Starting youtube service")
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping youtube service")
			return
		case <-ticker.C:
			log.Println("Searching videos for query", y.Query)
			y.searchAndStoreVideos()
		}
	}
}

func (y *YoutubeService) searchAndStoreVideos() {
}
