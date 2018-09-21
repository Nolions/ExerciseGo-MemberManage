package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRoute(r *gin.Engine) {
	r.Use(checkLogined())

	r.GET("/", index)
	r.GET("/login", login)
	r.POST("/login", attempt)
	r.GET("/result", result)
	r.GET("/logout", logout)

	// account mamager
	r.GET("/create", addAccount)
	r.POST("/create", sotreAccount)
	r.GET("/edit/:id", editAccount)
}

func reade(c *gin.Context, view string, data gin.H) {
	c.HTML(http.StatusOK, view, data)
}
