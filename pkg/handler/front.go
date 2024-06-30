package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}
