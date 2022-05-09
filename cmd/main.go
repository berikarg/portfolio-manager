package main

import (
	portfolio_manager "github.com/berikarg/portfolio-manager"
	"github.com/berikarg/portfolio-manager/pkg/handler"
	"github.com/berikarg/portfolio-manager/pkg/repository"
	"github.com/berikarg/portfolio-manager/pkg/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

func main() {
	logger, _ := zap.NewDevelopment()

	if err := initConfig(); err != nil {
		logger.Fatal("initialize configs", zap.Error(err))
	}

	db, err := initDB()
	if err != nil {
		logger.Fatal("initialize db", zap.Error(err))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, logger)

	srv := new(portfolio_manager.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Fatal("run server", zap.Error(err))
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDB() (*sqlx.DB, error) {
	addr := viper.GetString("database.addr")
	hostAndPort := strings.Split(addr, ":")
	if len(hostAndPort) != 2 {
		return nil, errors.New("invalid db address format")
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     hostAndPort[0],
		Port:     hostAndPort[1],
		Username: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
		SSLMode:  viper.GetString("database.ssl_mode"),
	})
	if err != nil {
		return nil, errors.Wrap(err, "create new db")
	}
	return db, nil
}
