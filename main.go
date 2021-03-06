package main

import (
	"fmt"
	"rest-api-courses/api/v1/Config"
	"rest-api-courses/api/v1/Models"
	"rest-api-courses/api/v1/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL())
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Course{}, &Models.CourseContent{})

	router := Routes.SetupRouter()

	router.Run()
}
