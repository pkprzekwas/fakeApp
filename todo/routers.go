package todo

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("todo/:id", Get)
	r.GET("/", List)
	r.POST("/", Create)
	r.PUT("/:id", Update)
	r.DELETE("/:id", Delete)
}

