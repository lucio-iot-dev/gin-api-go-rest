package main

import (
	"github.com/lucio-iot-dev/gin-api-go-rest/database"
	"github.com/lucio-iot-dev/gin-api-go-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()

}
