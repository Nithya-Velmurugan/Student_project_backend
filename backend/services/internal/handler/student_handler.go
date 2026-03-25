package handler

import "github.com/gin-gonic/gin"

func GetStudents(c *gin.Context) {
    c.JSON(200, gin.H{"data": "list of students"})
}

func RegisterStudentRoutes(r *gin.Engine) {
    r.GET("/students", GetStudents)
}
