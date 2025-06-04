package main

import (
	"fmt"
	"methodological-service/config"
	"methodological-service/handlers"
	"methodological-service/models"
	"methodological-service/repositories"
	routes "methodological-service/route"
	"methodological-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Discipline{}, &models.File{})

	repo := repositories.NewDisciplineRepository(db)
	service := services.NewDisciplineService(repo)
	handler := handlers.NewDisciplineHandler(service)

	r := gin.Default()
	routes.SetupRoutes(r, handler)

	r.Run(cfg.Port)
}
