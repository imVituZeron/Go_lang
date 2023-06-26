package main

import "github.com/gin-gonic/gin"

func showStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Vitor Santos",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", showStudents)
	r.Run(":5000")
}
