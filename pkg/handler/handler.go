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

	router.LoadHTMLGlob("pkg/templates/*")

	router.GET("/newURL", h.Index)
	router.GET("/c/:alias", h.RedirectByParams)

	api := router.Group("/api")
	{
		api.POST("/newURL", h.NewURL)
		api.GET("/aliasJSON", h.RedirectByJSON)
		api.GET("/aliasForm", h.RedirectByForm)

	}

	return router

}
