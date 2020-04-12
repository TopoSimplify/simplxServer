package main

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setRoutes(router *gin.Engine) {
	router.GET("/hello", server.hello)
	router.POST("/hello", server.simplify)
	router.POST("/simplify", server.simplify)
}

func (server *Server) hello(ctx *gin.Context) {
	ctx.JSON(Success, gin.H{"message": "hello"})
}

func (server *Server) simplify(ctx *gin.Context) {
	ctx.JSON(Success, gin.H{"message": "success"})
}
