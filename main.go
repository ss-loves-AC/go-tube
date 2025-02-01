package main

import (
	"go-tube/api"
	"go-tube/handler"
	"go-tube/storage"
)

func main() {

	db := storage.NewDB()
	cache := storage.NewCache()
	h := handler.NewHandler(db, cache)
	router := api.NewRouter(h)

	router.Run(":8080")
}