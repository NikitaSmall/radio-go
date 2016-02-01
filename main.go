package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.Static("/music", "./music")

	router.LoadHTMLGlob("templates/*")

	router.GET("/", indexHandler)
	router.Run(":3000")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
