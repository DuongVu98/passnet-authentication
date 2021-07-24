package main

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/rest"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config"
	app2 "github.com/DuongVu98/passnet-authentication/src/main/go/config/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	//err := godotenv.Load(getEnvFile())
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	app2.LoadEnv()

	serverPort := os.Getenv("SERVER_PORT")
	log.Printf("serverPort: %v", serverPort)

	app := echo.New()
	config.RunAppConfig()

	// Routes
	authRouting(app, "/auth")

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	// Start server
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", serverPort)))
}

func authRouting(app *echo.Echo, routerString string) {
	authGroup := app.Group(routerString)
	authGroup.POST("/register", rest.Register)
	authGroup.POST("/delete", rest.DeleteUser)
	authGroup.GET("/all", rest.GetAllUsers)
}

func getEnvFile() string {
	envFolder := "env/"
	env := os.Getenv("ENV")
	if env == "development" {
		return fmt.Sprintf("%v.env.dev", envFolder)
	}
	return ""
}
