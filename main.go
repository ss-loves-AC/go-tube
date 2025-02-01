package main

import (
	"context"
	"go-tube/api"
	"go-tube/handler"
	"go-tube/service"
	"go-tube/storage"
	"os"
	"os/signal"
	"time"
	"log"
)

func main() {

	log.Println("Starting server!!!")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	db := storage.NewDB()
	cache := storage.NewCache()

	apiKeys := []string{
		"API_KEY_1",
	}

	youtubeSvc := service.NewYoutubeService(db, cache, "golang", apiKeys)
	go youtubeSvc.Start(ctx, 10*time.Second)

	h := handler.NewHandler(db, cache)
	go func() {
		if err := api.Run(h, ":8080"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server!!!")

}
