package main

import (
	"avito_test_GO/config"
	"avito_test_GO/internal/app"
	"log"

	_ "avito_test_GO/docs"
)

// @title Тестовое задание на позицию стажёра-бэкендера от Avito Tech
// @description Микросервис для работы с балансом пользователей

// @host localhost:9000
// @BasePath /
func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
