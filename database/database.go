package database

import (
	"api-go-controle-investimentos/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComDB() {
	var usuario models.Usuario
	var investimento models.Investimento
	stringDeConexao := "host=localhost user=postgres password=root dbname=db_controle_investimentos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatal("erro ao se conectar com banco de dados")
	}

	if err := DB.AutoMigrate(&usuario, &investimento); err != nil {
		log.Fatal("Erro na migração")
	}
}
