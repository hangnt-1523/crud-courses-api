package main

import (
	"courses-api/controllers"
	"courses-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB()

	router := gin.Default()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.GET("/courses", controllers.GetAllCourses)
	router.GET("/courses/:id", controllers.FindCourse)
	router.POST("/courses", controllers.CreateCourse)
	router.PATCH("/courses/:id", controllers.UpdateCourse)
	router.DELETE("/courses/:id", controllers.DeleteCourse)

	router.Run()
}
