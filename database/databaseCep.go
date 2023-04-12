package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Erro ao carregar variaveis de ambiente para se conectar ao banco de dados", err)
	}

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	stringConection := fmt.Sprintf("user=%s dbname=%s password=%s  dbname=%s sslmode=disable", user, dbname, password, dbname)
	fmt.Println(stringConection)

	db, err := sql.Open("postgres", stringConection)
	if err != nil {
		log.Println("Erro ao se conectar ao banco de dados", err)

	}
	fmt.Println("Sucesso ao realizar conexão com banco de dados")

	return db

}
