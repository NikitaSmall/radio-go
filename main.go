package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/radio-go/socket"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.Static("/music", "./music")
	router.LoadHTMLGlob("templates/*")

	go socket.MusicHub.Run()

	router.GET("/", indexHandler)
	router.GET("/stream", hubHandler)

	router.Run(":3000")
}
