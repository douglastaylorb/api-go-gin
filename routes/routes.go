package routes

import (
	"github.com/douglastaylorb/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.GetAluno)
	r.POST("/alunos", controllers.CriaAluno)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditAluno)
	r.GET("/alunos/cpf/:cpf", controllers.SearchAlunoWithCPF)
	r.Run()
}
