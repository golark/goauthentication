package main

import (
	"github.com/gin-gonic/gin"
)

func headerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (app *application) loggerMiddleware() gin.HandlerFunc {
	return func(c * gin.Context) {
		app.infoLog.Printf(c.Request.Method)
	}
}

func setupGinRouter(app application) *gin.Engine {

	var router = gin.Default()

	router.Use(headerMiddleware())
	router.Use(app.loggerMiddleware())
	router.POST("/login", app.SignIn)
	router.POST("/task", app.CreateTask)
	router.POST("/authenticate", app.SignIn)

	return router
}
