package main

import (
	"log"

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
	); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	// Инициализация репозиториев
	categoryRepo := repository.NewCategoryRepository(db)
	subcategoryRepo := repository.NewSubcategoryRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	// Инициализация сервисов
	categoryService := services.NewCategoryService(categoryRepo)
	subcategoryService := services.NewSubcategoryService(subcategoryRepo)
	paymentService := services.NewPaymentService(paymentRepo)
	reviewService := services.NewReviewService(reviewRepo)

	// Инициализация хэндлеров
	categoryHandler := transport.NewCategoryHandler(categoryService)
	subcategoryHandler := transport.NewSubcategoryHandler(subcategoryService)
	paymentHandler := transport.NewPaymentHandler(paymentService)
	reviewHandler := transport.NewReviewHandler(reviewService)

	router := gin.Default()

	// Регистрация роутов
	categoryHandler.RegisterRoutes(router)
	subcategoryHandler.RegisterRoutes(router)
	paymentHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
