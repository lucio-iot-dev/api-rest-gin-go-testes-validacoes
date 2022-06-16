package main

import (
	"github.com/lucio-iot-dev/api_rest_gin_go_testes_validacoes/database"
	"github.com/lucio-iot-dev/api_rest_gin_go_testes_validacoes/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
