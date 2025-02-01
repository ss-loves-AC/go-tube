package api

import (
	"go-tube/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/videos", h.GetVideos)
	router.GET("/search", h.SearchVideos)

	return router
}

func Run(h *handler.Handler, addr string) error {
	router := NewRouter(h)
	return router.Run(addr)
}