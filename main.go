package main

import(
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
  c.JSON(200, gin.H{
		"id":"1",
		"nome":"Gui Lima",
	})
}

func main() {
	r := gin.Default()
	r.GET("/alunos", ExibeTodosAlunos )
	r.Run() // porta padão é 8080 do gin, caso queira subir em outra porta é so passar entre aspas dentro do parenteses do Run o :5000 por exemplo
}