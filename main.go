package main

import (
	"github.com/lucio-iot-dev/gin-api-go-rest/database"
	"github.com/lucio-iot-dev/gin-api-go-rest/models"
	"github.com/lucio-iot-dev/gin-api-go-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Lucio", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Beatriz", CPF: "11111111111", RG: "4800000000"},
	}

	routes.HandleRequests()

}
