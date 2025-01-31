package main

import (
	"go-tube/api"
	"go-tube/handler"
	"go-tube/database"
)

func main() {

	db := database.NewDB()
	cache := database.NewCache()
	h := handler.NewHandler(db, cache)
	router := api.NewRouter(h)

	router.Run(":8080")
}