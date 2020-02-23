package main

import (
	"github.com/andreylm/bmlabs/pkg/config"
	"github.com/andreylm/bmlabs/pkg/db/mongodb"
	"github.com/andreylm/bmlabs/pkg/logger"
	"github.com/andreylm/bmlabs/pkg/server"
)

func main() {
	logger.Init()
	configPaths := []string{
		"./", "./config/", "../config/, ../../config",
	}
	if err := config.Load("application.yaml", "yaml", configPaths...); err != nil {
		panic(err)
	}

	db := mongodb.NewStorage(
		config.ReadEnvString("DB_HOST"),
		config.ReadEnvInt("DB_PORT"),
		config.ReadEnvString("DB_NAME"),
		config.ReadEnvString("DB_USER"),
		config.ReadEnvString("DB_PASSWORD"),
	)

	srv := server.NewServer(config.ReadEnvString("APP_HOST"), config.ReadEnvInt("APP_PORT"), db)
	srv.Init()
	logger.Get().Panic(srv.Run())
}
