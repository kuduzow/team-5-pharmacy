package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/config"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"github.com/kuduzow/team-5-pharmacy/internal/repository"
	"github.com/kuduzow/team-5-pharmacy/internal/services"
	"github.com/kuduzow/team-5-pharmacy/internal/transport"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(
		&models.Category{},
		&models.Subcategory{},
		&models.Medicine{},
		&models.Payment{},
		&models.Review{},
		&models.User{},
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	// Инициализация репозиториев
	categoryRepo := repository.NewCategoryRepository(db)
	subcategoryRepo := repository.NewSubcategoryRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	userRepo := repository.NewUserRepository(db)
	medicineRepo := repository.NewMedicinesRepository(db)

	// Инициализация сервисов
	categoryService := services.NewCategoryService(categoryRepo)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	reviewService := services.NewReviewService(reviewRepo)
	userService := services.NewUserService(userRepo)
	medicineService := services.NewMedicineService(medicineRepo)

	router := gin.Default()

	transport.RegisterRoutes(router, categoryService, subcategoryService, paymentService, reviewService, userService)
	transport.NewHandlerMedicine(medicineService).RegisterRoutes(router)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // поменяй на "8081", если 8080 занят
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
