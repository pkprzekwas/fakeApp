package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkprzekwas/fakeApp/common"
	"github.com/pkprzekwas/fakeApp/config"
	"github.com/pkprzekwas/fakeApp/todo"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found.")
	}
	conf := config.New()
	common.Init(&conf.DBConfig)
	todo.AutoMigrate()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1/")
	todo.RegisterRoutes(v1)

	return router
}

func main() {
	engine := setupRouter()
	if err := engine.Run(); err != nil {
		panic(err)
	}
}
