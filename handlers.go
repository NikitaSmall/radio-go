package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/radio-go/socket"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func currentTrackHandler(c *gin.Context) {
	trackInfo, _ := socket.MusicHub.TrackInfo()

	c.JSON(http.StatusOK, trackInfo.Duration())
}
