package handler

import (
	"net/http"
	"shorter/pkg/gen"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Shorting(c *gin.Context) {
	str := gen.GenerateAlias(5)
	c.HTML(http.StatusOK, "form.html", gin.H{
		"str": str,
	})
}

func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}
