package main

import(
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run() // porta padão é 8080 do gin, caso queira subir em outra porta é so passar entre aspas dentro do parenteses do Run o :5000 por exemplo
}