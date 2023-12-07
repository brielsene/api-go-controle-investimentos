package main

import (
	"api-go-controle-investimentos/database"
	"api-go-controle-investimentos/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()

}
