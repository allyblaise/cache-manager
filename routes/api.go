package routes

import (
	"cache-manager/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1/")

	v1.GET("service-status", controller.ServiceStatusCheck)
	v1.GET("deleteCache/:tag", controller.DeleteCache)
	v1.POST("addCache", controller.AddCache)
	return r
}
