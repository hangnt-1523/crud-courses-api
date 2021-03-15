package Routes

import (
  "rest-api-courses/api/v1/Controllers"
  "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
  router := gin.Default()

  grp1 := router.Group("/courses-api")
  {
    grp1.GET("/courses", Controllers.GetCourses)
    grp1.GET("/courses/:id", Controllers.GetCourseByID)
    grp1.POST("/courses", Controllers.CreateCourse)
    grp1.PATCH("/courses/:id", Controllers.UpdateCourse)
    grp1.DELETE("/courses/:id", Controllers.DeleteCourse)
  }

  return router
}
