package main

import (
	"context"
	"go-tube/api"
	"go-tube/handler"
	"go-tube/service"
	"go-tube/storage"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	log.Println("Starting server!!!")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	db , err := storage.NewDB()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}	
	cache := storage.NewCache()

	apiKeys := []string{
		"AIzaSyA8Dlbjo9j9Hcg4UH1Hjh8InuZ0BexKC10",
	}
    
	predfinedQuery := "cricket"

	youtubeSvc := service.NewYoutubeService(db, cache, predfinedQuery , apiKeys)
	go youtubeSvc.Start(ctx, 10*time.Second , cancel)

	h := handler.NewHandler(db, cache)
	go func() {
		if err := api.Run(h, ":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("Shutting down server!!!")
	case <-signalChan:
		log.Println("Received interrupt signal")
		cancel()
	}
	
}
