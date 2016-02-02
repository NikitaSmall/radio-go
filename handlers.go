package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nikitasmall/radio-go/config"
	"github.com/nikitasmall/radio-go/model"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func startTrackHandler(c *gin.Context) {
	db := config.GetConnection()
	defer db.Close()

	var song model.Song
	db.Order("id desc").First(&song)

	c.JSON(http.StatusOK, song)
}

func nextTrackHandler(c *gin.Context) {
	id := c.Param("id")

	db := config.GetConnection()
	defer db.Close()

	var song model.Song
	db.Order("id desc").Where("id < ?", id).First(&song)

	c.JSON(http.StatusOK, song)
}
