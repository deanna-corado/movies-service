package main

import (
	"log"
	"movies-service/config"
	"movies-service/controllers"
	_ "movies-service/docs"
	"movies-service/migrations"
	"movies-service/repositories"
	"movies-service/routes"
	"movies-service/services"
	"os"

	"github.com/gin-gonic/gin"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/joho/godotenv"
)

// @title Movies API
// @version 1.0
// @description Server for managing movies
// @host localhost:5002
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	m := gormigrate.New(
		config.DB,
		gormigrate.DefaultOptions,
		migrations.GetMigrations(),
	)

	//MIGRATE ALL LISTED MIGRATIONS
	if err := m.Migrate(); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations applied successfully")

	//ROLLBACK LAST MIGRATION ONLY

	// if err := m.RollbackLast(); err != nil {
	// 	log.Fatal(err)
	// }

	//ROLLBACK SPECIFIC MIGRATION

	// if err := m.RollbackTo("00003_createCredentialTable"); err != nil {
	// 	log.Fatal(err)
	// }

	r := gin.Default()

	movieRepo := repositories.NewMovieRepository(config.DB)
	movieService := services.NewMovieService(movieRepo)
	movieController := controllers.NewMovieController(movieService)
	credController := controllers.NewCredentialController(config.DB)

	routes.RegisterRoutes(r, movieController, credController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5002"
	}

	log.Println("Server running on port", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
