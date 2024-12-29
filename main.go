package main

import (
	"github.com/douglastaylorb/api-go-gin/database"
	"github.com/douglastaylorb/api-go-gin/models"
	"github.com/douglastaylorb/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Douglas", CPF: "11145201471", RG: "5000000"},
		{Nome: "João", CPF: "00045201471", RG: "5000000"},
		{Nome: "Maria", CPF: "22245201471", RG: "5000000"},
	}
	routes.HandleRequests()
}
