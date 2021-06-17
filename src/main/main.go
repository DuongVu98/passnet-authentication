package main

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/rest"
	"github.com/DuongVu98/passnet-authentication/src/main/go/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(getEnvFile())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	log.Printf("serverPort: %v", serverPort)

	app := echo.New()

	// wire beans (manually)
	log.Printf("before run config")
	config.RunAppConfig()
	log.Printf("after run config")

	// Routes
	app.GET("/", rest.Hello)
	homeRouting(app, "/test")
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

func homeRouting(app *echo.Echo, routerString string) {
	homeGroup := app.Group(routerString)

	// Methods
	homeGroup.GET("/", rest.HomePage)
	homeGroup.GET("/json", rest.JsonResponseSample)
	homeGroup.GET("/test-grpc", rest.TestGrpcMessage)
}
func authRouting(app *echo.Echo, routerString string) {
	authGroup := app.Group(routerString)
	authGroup.POST("/register", rest.Register)
}

func getEnvFile() string {
	envFolder := "env/"
	env := os.Getenv("ENV")
	if env == "development" {
		return fmt.Sprintf("%v.env.dev", envFolder)
	} else if env == "vagrant" {
		return fmt.Sprintf("%v.env.vagrant", envFolder)
	} else if env == "production" {
		return fmt.Sprintf("%v.env.prod", envFolder)
	}
	return ""
}
