package main // import "server"

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
  "server/handler"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	r.POST("/images",handler.Upload)
	r.DELETE("/images/:uuid", handler.Delete)
	r.Run(":8888")
}