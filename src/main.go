package main

import (
	"github.com/adlighozian/go-belajar/controllers"

	"github.com/adlighozian/go-belajar/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/products", productController.Create)
	r.PUT("/api/products/:id", productController.Update)
	r.DELETE("/api/products", productController.Delete)

	r.Run()
}
