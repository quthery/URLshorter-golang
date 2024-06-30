package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	sch "shorter/pkg/schemas"

	"github.com/gin-gonic/gin"
)

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

func (h *Handler) NewURL(c *gin.Context) {

	var json sch.URL

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.database.SaveURL(json.URL, json.Alias)
	if err != nil {
		fmt.Print("Error:", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	slog.Info("Url saved at id: ", slog.Int64("id", id))
	c.JSON(200, gin.H{
		"id":    id,
		"url":   json.URL,
		"alias": json.Alias,
	})

}

func (h *Handler) RedirectByJSON(c *gin.Context) {
	var alias sch.AliasGet

	if err := c.ShouldBindJSON(&alias); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := h.database.GetURL(alias.Alias)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	slog.Info("response url succes")
	c.JSON(200, gin.H{
		"url": url,
	})

}

func (h *Handler) RedirectByForm(c *gin.Context) {
	var alias sch.AliasGet

	if err := c.ShouldBind(&alias); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := h.database.GetURL(alias.Alias)
	if err != nil {
		c.JSON(http.StatusBadRequest, "wrong alias!")
		return
	}
	c.Redirect(301, url)
	slog.Info("response url succes")
	c.JSON(200, gin.H{
		"url": url,
	})

}

func (h *Handler) RedirectByParams(c *gin.Context) {
	var alias sch.AliasGet

	if err := c.ShouldBindQuery(&alias); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url, err := h.database.GetURL(alias.Alias)
	if err != nil {
		c.JSON(http.StatusBadRequest, "wrong alias!")
		return
	}
	c.Redirect(301, url)
	slog.Info("response url succes")
	c.JSON(200, gin.H{
		"url": url,
	})

}
