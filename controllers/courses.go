package controllers

import (
	"net/http"

	"courses-api/models"

	"github.com/gin-gonic/gin"
)

type CreateCourseInput struct {
	Title string `json:"title" binding:"required"`
}

type UpdateCourseInput struct {
	Title string `json:"title"`
}

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	models.DB.Find(&courses)

	c.JSON(http.StatusOK, gin.H{"data": courses})
}

func FindCourse(c *gin.Context) {
	var course models.Course
	if err := models.DB.Where("id = ?", c.Param("id")).First(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func CreateCourse(c *gin.Context) {
	var input CreateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Course{Title: input.Title}
	models.DB.Create(&course)

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	if err := models.DB.Where("id = ?", c.Param("id")).First(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateCourseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&course).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": course})
}

func DeleteCourse(c *gin.Context) {
	var course models.Course
	if err := models.DB.Where("id = ?", c.Param("id")).First(&course).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&course)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
