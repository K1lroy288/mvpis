package main

import (
	"educational-service/config"
	"educational-service/handlers"
	"educational-service/models"
	"educational-service/repositories"
	"educational-service/routes"
	"educational-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к PostgreSQL через GORM
	dsn := "host=localhost user=postgres password=3321 dbname=mvpb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Автомиграция
	db.AutoMigrate(&models.Publication{}, &models.Conference{}, &models.Thesis{}, &models.Supervisor{})

	// Создание репозитория, сервиса и обработчика
	repo := repositories.NewResearchRepository(db)
	service := services.NewResearchService(repo)
	handler := handlers.NewResearchHandler(service)

	// Настройка маршрутов
	r := gin.Default()
	routes.SetupResearchRoutes(r, handler)

	// Запуск сервера
	r.Run(cfg.Port)
}