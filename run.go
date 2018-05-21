package main

import (
	"gopkg.in/gin-gonic/gin.v1"
)

func (server *Server) Run() {
	var router = gin.Default()
	gin.SetMode(server.Mode)
	server.setRoutes(router)
	router.Run(server.Address)
}
