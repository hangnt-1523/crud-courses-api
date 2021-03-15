package Controllers
import (
  "rest-api-courses/api/v1/Models"
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
  var courses []Models.Course
  err := Models.GetAllCourses(&courses)
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, courses)
  }
}

func CreateCourse(c *gin.Context) {
  var course Models.Course
  c.BindJSON(&course)
  err := Models.CreateCourse(&course)
  if err != nil {
    fmt.Println(err.Error())
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, course)
  }
}

func GetCourseByID(c *gin.Context) {
  id := c.Params.ByName("id")
  var course Models.Course
  err := Models.GetCourseByID(&course, id)
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, course)
  }
}

func UpdateCourse(c *gin.Context) {
  var course Models.Course
  id := c.Params.ByName("id")
  err := Models.GetCourseByID(&course, id)
  if err != nil {
    c.JSON(http.StatusNotFound, course)
  }
  c.BindJSON(&course)
  err = Models.UpdateCourse(&course, id)
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, course)
  }
}

func DeleteCourse(c *gin.Context) {
  var course Models.Course
  id := c.Params.ByName("id")
  err := Models.DeleteCourse(&course, id)
  if err != nil {
    c.AbortWithStatus(http.StatusNotFound)
  } else {
    c.JSON(http.StatusOK, gin.H{"Course" + id: "is deleted"})
  }
}
