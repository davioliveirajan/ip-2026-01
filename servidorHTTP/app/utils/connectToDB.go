package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectToDB() {
	err := godotenv.Load("servidorHTTP/.env")
	if err != nil {
		err = godotenv.Load(".env")
	}
	if err != nil {
		err = godotenv.Load("../.env")
	}
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if dbname == "" {
		dbname = user
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao verificar a conexao com o banco de dados: %v", err)
	}

	fmt.Println("Conexao com o banco de dados estabelecida com sucesso!")
}
