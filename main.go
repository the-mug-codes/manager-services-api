package main

import (
	"os"
	"strconv"

	"github.com/kodit-tecnologia/service-manager/docs"
	router "github.com/kodit-tecnologia/service-manager/presenters/routes"
	"github.com/kodit-tecnologia/service-manager/repositories"
	"github.com/the-mug-codes/adapters-service-api/auth"
	cron "github.com/the-mug-codes/adapters-service-api/cron"
	database "github.com/the-mug-codes/adapters-service-api/database"
	server "github.com/the-mug-codes/adapters-service-api/server"
	utils "github.com/the-mug-codes/adapters-service-api/utils"
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
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	utils.CreateTemporaryDirectory("tmp")
	utils.LoadEnvironmentVariables(os.Getenv("MODE"))
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 80
	}
	databasePort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		port = 5432
	}
	auth.Connect(auth.AuthData{
		Host:         os.Getenv("AUTH_HOST"),
		Realm:        os.Getenv("AUTH_REALM"),
		ClientID:     os.Getenv("AUTH_CLIENT"),
		ClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		PublicKey:    os.Getenv("AUTH_PUBLIC_KEY"),
	})
	database.Start(database.PostgresData{
		Port:         databasePort,
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Host:         os.Getenv("DATABASE_HOST"),
		Username:     os.Getenv("DATABASE_USERNAME"),
		Password:     os.Getenv("DATABASE_PASSWORD"),
		Timezone:     os.Getenv("DATABASE_TIMEZONE"),
	}, os.Getenv("MODE"))
	database.RunMigrations(repositories.Migrations)
	cron.AutomaticJobs([]cron.CronAutomaticJobs{})
	server.Start(os.Getenv("INSTANCE"), os.Getenv("NAME"), port, "0.1.0", router.Routes, nil)
}
