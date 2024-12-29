package controllers

import (
	"github.com/douglastaylorb/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}
