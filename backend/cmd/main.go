package main

import (
	"log"
	"personal-finance-app/internal/platform/database"
	"personal-finance-app/internal/platform/router"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Получаем sqlDB и обрабатываем ошибку
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	r := router.SetupRouter(db)
	r.Run(":8080")
}
