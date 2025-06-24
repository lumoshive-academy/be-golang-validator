package main

import (
	"fmt"
	"go-23/database"
	"go-23/handler"
	"go-23/repository"
	"go-23/router"
	"go-23/service"
	"go-23/utils"
	"log"
	"net/http"

	"go.uber.org/zap"
)

func main() {
	// init logger
	logger, err := utils.InitLogger("./logs/app-", true)
	if err != nil {
		log.Fatal("can't init logger %w", zap.Error(err))
	}

	//Init db
	db, err := database.InitDB()
	if err != nil {
		logger.Fatal("can't connect to database ", zap.Error(err))
	}

	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service)

	r := router.NewRouter(handler)

	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal("can't run service", zap.Error(err))
	}
}
