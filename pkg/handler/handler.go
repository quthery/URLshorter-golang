package handler

import (
	sqlite "shorter/pkg/sqlite"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	database *sqlite.Storage
}

func NewHandler(database *sqlite.Storage) *Handler {
	return &Handler{database: database}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/newURL", h.NewURL)
		api.GET("/aliasJSON", h.RedirectByJSON)
		api.GET("/aliasForm", h.RedirectByForm)
		api.GET("/alias", h.RedirectByParams)
	}

	return router

}
