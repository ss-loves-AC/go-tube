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
}

func (h *Handler) SearchVideos(c *gin.Context) {
}

func NewHandler(db storage.DB , cache storage.Cache) *Handler {
	return &Handler{
		DB: db,
		Cache: cache,
	}
}
