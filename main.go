package main

import (
	"github.com/douglastaylorb/api-go-gin/database"
	"github.com/douglastaylorb/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
