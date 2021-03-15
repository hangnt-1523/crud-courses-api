package Models

import (
  "rest-api-courses/api/v1/Config"
  "fmt"
_ "github.com/go-sql-driver/mysql"
)

func GetAllCourses(courses *[]Course) (err error) {
  if err = Config.DB.Find(courses).Error; err != nil {
    return err
  }
  return nil
}

func CreateCourse(course *Course) (err error) {
  if err = Config.DB.Create(course).Error; err != nil {
    return err
  }
  return nil
}

func GetCourseByID(course *Course, id string) (err error) {
  if err = Config.DB.Where("id = ?", id).First(course).Error; err != nil {
    return err
  }
  return nil
}

func UpdateCourse(course *Course, id string) (err error) {
  fmt.Println(course)
  Config.DB.Save(course)
  return nil
}

func DeleteCourse(course *Course, id string) (err error) {
  Config.DB.Where("id = ?", id).Delete(course)
  return nil
}