package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Define a URL de conexão com o banco de dados MySQL
	dbURL := "mysql://root:social@localhost:9030/social?charset=utf8mb4&parseTime=true&loc=Local"

	// Define o caminho para as migrações
	migrationsPath := "migrations"

	// Cria um novo objeto Migrate
	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	// Aplica as migrações
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	fmt.Println("Migrações aplicadas com sucesso!")
}
