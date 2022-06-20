package main

import (
	"github.com/lucio-iot-dev/api-rest-gin-go-testes-validacoes/database"
	"github.com/lucio-iot-dev/api-rest-gin-go-testes-validacoes/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
