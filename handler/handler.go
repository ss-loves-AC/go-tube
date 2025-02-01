package handler

import (
	"go-tube/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB    storage.DB
	Cache storage.Cache
}

func (h *Handler) GetVideos(c *gin.Context) {
	page, limit := getPaginatedParams(c)
	videos, err := h.DB.GetVideos(c, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  videos,
	})
}

func (h *Handler) SearchVideos(c *gin.Context) {
	query := c.Query("query")
	page, limit := getPaginatedParams(c)

	if query == "" {
		c.JSON(400, gin.H{"error": "query parameter is required"})
		return
	}
	videos, err := h.DB.SearchVideos(c, query, page, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"data":  videos,
	})
}

func NewHandler(db storage.DB, cache storage.Cache) *Handler {
	return &Handler{
		DB:    db,
		Cache: cache,
	}
}

func getPaginatedParams(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return page, limit
}
