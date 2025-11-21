package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/config"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
)

func main() {
	db := config.SetUpDatabaseConnection()
	if db == nil {
		log.Fatal("Не удалось подключиться к базе данных")
	}
	fmt.Println(db) // временно

	// Выполняем миграции моделей
	// AutoMigrate (создаст tables)
	if err := db.AutoMigrate(&models.Review{}); err != nil {
		log.Fatal("migrate:", err)
	}

	fmt.Println("DB migrated")

	router := gin.Default()
	// reviewHandler := handlers.NewReviewHandler(db)

	// router.POST("/reviews", reviewHandler.Create)
	// router.GET("/medicines/:id/reviews", reviewHandler.GetByMedicineID)
	// router.PATCH("/reviews/:id", reviewHandler.Update)
	// router.DELETE("/reviews/:id", reviewHandler.Delete)

	// регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
