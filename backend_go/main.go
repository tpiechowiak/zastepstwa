package main

import (
	"github.com/gin-gonic/gin"
	"zsk.poznan.pl/backend/handlers/date"
	"zsk.poznan.pl/backend/handlers/subs"
	"zsk.poznan.pl/backend/handlers/teacher"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		v1.GET("teacher", teacher.GetTeachers)
		v1.GET("teacher/:name", teacher.GetTeacherByName)
		v1.GET("teacher/all", subs.GetAllSubs)
		v1.GET("date", date.GetDate)
	}

	r.Run()
}
