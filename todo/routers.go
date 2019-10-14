package todo

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("todo/:id", Get)
	r.GET("todo/", List)
	r.POST("todo/", Create)
	r.PUT("todo/:id", Update)
	r.DELETE("todo/:id", Delete)
}

