package main

import (
	"log"
	"tugas-oti/config"
	"tugas-oti/docs"
	"tugas-oti/routes"
	"tugas-oti/utils"

	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	docs.SwaggerInfo.Title = "Swagger Event Enrollment API"
	docs.SwaggerInfo.Description = "This is a event enrollment server."
	docs.SwaggerInfo.Version = "1.0"

	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
