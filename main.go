package main

import (
	"elib_v2/controllers"
	"elib_v2/initializers"
	"elib_v2/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//auth
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	//crud buku
	r.GET("/books", controllers.Books)
	r.GET("/books/:id", controllers.Viewbookid)
	r.POST("/addbook", middleware.RequireAuth, controllers.AddBooks)
	r.PUT("/books/:id", middleware.RequireAuth, controllers.UpdateBook)
	r.DELETE("/books/:id", middleware.RequireAuth, controllers.DeleteBook)

	r.Run()
}
