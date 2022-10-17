package main

import (
	"mygram/database"
	"mygram/routers"
	"mygram/docs"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/
func main() {
	docs.SwaggerInfo.Title = "MyGram API"
    docs.SwaggerInfo.Description = "This is a MyGram API server."
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}