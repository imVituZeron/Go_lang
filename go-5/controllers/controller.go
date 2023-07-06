package controllers

import (
	"net/http"
	"pack/api-gin/database"
	"pack/api-gin/models"
	M "pack/api-gin/models"

	"github.com/gin-gonic/gin"
)

func ShowStudents(c *gin.Context) {
	c.JSON(200, M.Alunos)
}

func Welcome(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + " tudo beleza?.",
	})
}

func CreateNewStudent(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}
