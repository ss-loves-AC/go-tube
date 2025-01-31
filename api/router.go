package api

import (
	"go-tube/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/videos", h.GetVideos)
	router.GET("/search", h.SearchVideos)

	return router
}