package main

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) Run() {
	gin.SetMode(server.Mode)
	var router = gin.Default()
	server.setRoutes(router)
	router.Run(server.Address)
}
