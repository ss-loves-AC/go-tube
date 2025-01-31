package handler

import (
	"go-tube/database"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB database.DB
	Cache database.Cache
}

func (h *Handler) GetVideos(c *gin.Context) {

}

func (h *Handler) SearchVideos(c *gin.Context) {
}

func NewHandler(db database.DB , cache database.Cache) *Handler {
	return &Handler{
		DB: db,
		Cache: cache,
	}
}
