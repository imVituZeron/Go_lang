package routes

import (
	C "pack/api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", C.ShowStudents)
	r.GET("/:nome", C.Welcome)
	r.POST("/alunos", C.CreateNewStudent)
	r.Run(":5000")
}
