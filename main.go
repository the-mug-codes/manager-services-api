package main

import (
	"os"
	"strconv"

	cron "github.com/the-mug-codes/adapters-service-api/cron"
	database "github.com/the-mug-codes/adapters-service-api/database"
	server "github.com/the-mug-codes/adapters-service-api/server"
	utils "github.com/the-mug-codes/adapters-service-api/utils"
	"github.com/the-mug-codes/service-manager-api/adapters/websocket"
	router "github.com/the-mug-codes/service-manager-api/presenters/routes"
	"github.com/the-mug-codes/service-manager-api/repositories"
)

// @title						The Mug Codes
// @version					1.0
// @description				Welcome to the official documentation for the The Mug Codes manager. This guide provides comprehensive information to integrate and utilize our API effectively. It includes details on authentication, endpoints, rate limits, error handling, best practices, and support. Whether you're a beginner or an experienced developer, this documentation serves as your roadmap to leverage our services seamlessly.
// @BasePath					/v1
// @host						localhost:9000
// @schemes http https
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	utils.CreateTemporaryDirectory("tmp")
	utils.LoadEnvironmentVariables(os.Getenv("MODE"))
	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 80
	}
	databasePort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		port = 5432
	}
	database.Start(database.PostgresData{
		Port:         databasePort,
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Host:         os.Getenv("DATABASE_HOST"),
		Username:     os.Getenv("DATABASE_USERNAME"),
		Password:     os.Getenv("DATABASE_PASSWORD"),
		SSLMode:      "require",
		Timezone:     os.Getenv("DATABASE_TIMEZONE"),
	}, os.Getenv("MODE"))
	if len(os.Getenv("MIGRATE")) != 0 {
		database.RunMigrations(repositories.Migrations)
	}
	websocket.StartWebsocketChat()
	cron.AutomaticJobs([]cron.CronAutomaticJobs{})
	server.Start(os.Getenv("INSTANCE"), os.Getenv("APP_NAME"), port, "0.1.0", router.Routes, nil)
}
