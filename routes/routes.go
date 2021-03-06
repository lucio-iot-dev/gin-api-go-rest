package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucio-iot-dev/gin-api-go-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.PUT("/alunos/:id", controllers.AtualizaAluno)

	r.Run() // porta padão é 8080 do gin, caso queira subir em outra porta é so passar entre aspas dentro do parenteses do Run o :5000 por exemplo
}
