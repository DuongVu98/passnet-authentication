package main

import (
	"fmt"
	"github.com/DuongVu98/passnet-authentication/src/main/go/adapter/midlewares"
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
	userRouting(app, "/user")

	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

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
	authGroup.Use(midlewares.GetBeanMiddlewareProcess)

	// Methods
	authGroup.POST("/login", rest.Login)
	authGroup.POST("/signup", rest.SignUp)
}
func userRouting(app *echo.Echo, routerString string) {
	userGroup := app.Group(routerString)
	userGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("AUTH_SECRET")),
		TokenLookup: os.Getenv("AUTH_TOKEN_LOOKUP"),
		ContextKey:  os.Getenv("AUTH_CONTEXT_KEY"),
		AuthScheme:  os.Getenv("AUTH_SCHEME"),
	}), middleware.KeyAuth(midlewares.CheckUserExistMiddleware))

	// Methods
	userGroup.GET("/user-info/:uid", rest.UserInfo)
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
