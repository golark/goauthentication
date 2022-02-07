package main

import "github.com/gin-gonic/gin"

func setupGinRouter(app application) *gin.Engine {

	var router = gin.Default()

	router.POST("/login", app.SignIn)
	router.POST("/todo", app.CreateTodo)
	router.POST("/logout", app.Logout)
	router.POST("/refresh", app.Refresh)

	return router
}