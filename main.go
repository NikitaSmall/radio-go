package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.Static("/music", "./music")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)
	router.GET("/start", startTrackHandler)
	router.GET("/next/:id", nextTrackHandler)

	router.Run(":3000")
}
