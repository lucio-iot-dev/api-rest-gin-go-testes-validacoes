package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucio-iot-dev/api_rest_gin_go_testes_validacoes/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:nome", controllers.Saudacoes)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run()
}
