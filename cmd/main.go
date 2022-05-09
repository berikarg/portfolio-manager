package main

import (
	portfolio_manager "github.com/berikarg/portfolio-manager"
	"github.com/berikarg/portfolio-manager/pkg/handler"
	"github.com/berikarg/portfolio-manager/pkg/repository"
	"github.com/berikarg/portfolio-manager/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("initialize configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(portfolio_manager.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
