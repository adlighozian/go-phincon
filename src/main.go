package main

import (
	"github.com/adlighozian/go-belajar/controllers/authController"
	"github.com/adlighozian/go-belajar/controllers/companyController"
	"github.com/adlighozian/go-belajar/controllers/homeController"
	"github.com/adlighozian/go-belajar/controllers/logsLoginController"
	"github.com/adlighozian/go-belajar/controllers/logsPresentController"
	"github.com/adlighozian/go-belajar/controllers/productController"
	"github.com/adlighozian/go-belajar/controllers/userController"
	"github.com/adlighozian/go-belajar/middlewares"
	"github.com/adlighozian/go-belajar/models"
	"github.com/gin-gonic/gin"
)

func main() {
	phincon()
	belajar()
}

func phincon() {
	r := gin.New()
	models.ConnectDatabase()

	authorized := r.Group("/api")
	authorized.Use(middlewares.JWTMiddleware)
	{
		// USER START
		authorized.GET("/user", userController.Index)
		authorized.GET("/user/:id", userController.Show)
		authorized.PUT("/user/:id", userController.Update)
		authorized.DELETE("/user", userController.Delete)
		// USER END
		// COMPANY START
		authorized.GET("/company", companyController.Index)
		authorized.GET("/company/:id", companyController.Show)
		authorized.POST("/company", companyController.Create)
		authorized.PUT("/company/:id", companyController.Update)
		authorized.DELETE("/company", companyController.Delete)
		// COMPANY GET
		// LOGS_LOGIN START
		authorized.GET("/loglogin", logsLoginController.Index)
		authorized.GET("/loglogin/:id", logsLoginController.Show)
		authorized.POST("/loglogin", logsLoginController.Create)
		authorized.PUT("/loglogin/:id", logsLoginController.Update)
		authorized.DELETE("/loglogin", logsLoginController.Delete)
		// LOGS_LOGIN END
		// LOGS_PRESENT START
		authorized.GET("/logpresent", logsPresentController.Index)
		authorized.GET("/logpresent/:id", logsPresentController.Show)
		authorized.POST("/logpresent", logsPresentController.Create)
		authorized.PUT("/logpresent/:id", logsPresentController.Update)
		authorized.DELETE("/logpresent", logsPresentController.Delete)
		// LOGS_PRESENT END
		// HOME START
		authorized.GET("/home", homeController.Index)
		authorized.GET("/home/:id", homeController.Show)
		authorized.POST("/home", homeController.Create)
		authorized.PUT("/home/:id", homeController.Update)
		authorized.DELETE("/home", homeController.Delete)
		// HOME END
	}
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	r.GET("/logout", authController.Logout)
	r.Run()
}

func belajar() {
	r := gin.New()
	r.GET("/api/products", productController.Index)
	r.GET("/api/products/:id", productController.Show)
	r.POST("/api/products", productController.Create)
	r.PUT("/api/products/:id", productController.Update)
	r.DELETE("/api/products", productController.Delete)
	r.Run()
}
