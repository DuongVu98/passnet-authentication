package main

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/adapter/midlewares"
	"github.com/DuongVu98/passnet-authentication/src/main/adapter/rest"
	"github.com/DuongVu98/passnet-authentication/src/main/config"
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

	app := echo.New()

	// wire beans (manually)
	log.Printf("before run config")
	config.RunAppConfig()
	log.Printf("after run config")

	// Routes
	app.GET("/", rest.Hello)
	homeRouting(app)
	userRouting(app)

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())


	// Start server
	app.Logger.Fatal(app.Start(fmt.Sprintf(":%s", serverPort)))
}

func homeRouting(homeRouter *echo.Echo) {
	homeRouter.GET("/page", rest.HomePage)
	homeRouter.GET("/page/json", rest.JsonResponseSample)
	homeRouter.GET("/test/user", rest.UserRetrieve)
}
func userRouting(userRouter *echo.Echo) {
	userRouter.POST("/user/login", rest.Login, midlewares.GetBeanMiddlewareProcess)
	userRouter.POST("/user/signup", rest.SignUp, midlewares.GetBeanMiddlewareProcess)
}

func getEnvFile() string {
	env := os.Getenv("ENV")
	if env == "development" {
		return ".env.dev"
	} else if env == "production" {
		return ".env.prod"
	}
	return ""
}
