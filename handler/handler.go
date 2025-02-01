package handler

import (
	"go-tube/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB storage.DB
	Cache storage.Cache
}

func (h *Handler) GetVideos(c *gin.Context) {
	videos, err := h.DB.GetVideos(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, videos)
}

func (h *Handler) SearchVideos(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(400, gin.H{"error": "query parameter is required"})
		return
	}
	videos, err := h.DB.SearchVideos(c, query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, videos)
}

func NewHandler(db storage.DB , cache storage.Cache) *Handler {
	return &Handler{
		DB: db,
		Cache: cache,
	}
}
