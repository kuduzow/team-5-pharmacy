package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/config"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Subcategory{},
		&models.Medicine{},
		&models.Payment{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

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
