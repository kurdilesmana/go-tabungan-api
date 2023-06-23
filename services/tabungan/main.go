package main

import (
	"github.com/kurdilesmana/go-tabungan-api/pkg/logging"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/api"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/app"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/datastore"
	"github.com/kurdilesmana/go-tabungan-api/services/tabungan/server"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("././")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// APP configuration
	APP_NAME := viper.GetString("APP_NAME")
	APP_PORT := viper.GetString("APP_PORT")

	// Database configuration
	DB_HOST := viper.GetString("DB_HOST")
	DB_PORT := viper.GetInt("DB_PORT")
	DB_USER := viper.GetString("DB_USER")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	DB_DATABASE := viper.GetString("DB_DATABASE")

	// Dependency injection
	logger := logging.NewLogger(APP_NAME)
	ds := datastore.Init(DB_USER, DB_PASSWORD, DB_DATABASE, DB_HOST, DB_PORT, logger)
	app := app.InitApplication(ds, logger)
	api := api.InitTabunganAPI(app, logger)

	// start service API
	server := server.InitServer(*api, logger)
	server.Start(APP_PORT)
}
