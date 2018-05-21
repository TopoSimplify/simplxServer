package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func (server *Server) setRoutes(router *gin.Engine) {
	router.GET("/hello", server.hello)
	router.POST("/hello", server.simplify)
	router.POST("/simplify", server.simplify)
}

func (s *Server) hello(ctx *gin.Context) {
	ctx.JSON(Success, gin.H{"message": "hello"})
}

func (s *Server) simplify(ctx *gin.Context) {
	ctx.JSON(Success, gin.H{"message": "success"})
}
