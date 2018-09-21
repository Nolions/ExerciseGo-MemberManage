package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// use defaut middleware, include(Logger and Recovery)
	router := gin.Default()

	// Load templates
	router.LoadHTMLGlob("views/*")

	// setting routies
	setRoute(router)

	router.Run()
}

func index(c *gin.Context) {
	data := gin.H{
		"title": "index",
	}

	reade(c, "index.html", data)
}
