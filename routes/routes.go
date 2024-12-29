package routes

import (
	"github.com/douglastaylorb/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.GetAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
