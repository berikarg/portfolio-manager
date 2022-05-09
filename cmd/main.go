package main

import (
	portfolio_manager "github.com/berikarg/portfolio-manager"
	"github.com/berikarg/portfolio-manager/pkg/handler"
	"github.com/berikarg/portfolio-manager/pkg/repository"
	"github.com/berikarg/portfolio-manager/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(portfolio_manager.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
