package api

import (
	"clean-architecture/api/routers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	routers.SetupRoutes(router)
	return router

}
