package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kuduzow/team-5-pharmacy/internal/config"
	"github.com/kuduzow/team-5-pharmacy/internal/config/models"
)

func main() {
	db := config.SetUpDatabaseConnection()

	fmt.Println(db) // временно

	if err := db.AutoMigrate(&models.Payment{}); err != nil{
		log.Fatalf("не удалось выполнить миграции:%v",err)
	}


	// Выполняем миграции моделей
	if err := db.AutoMigrate(&models.Medicine{}); err != nil{
		log.Fatalf("не удалось сделать миграции")
	}
	

	router := gin.Default()

	// регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
