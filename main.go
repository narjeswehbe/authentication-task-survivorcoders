package main

import (
	"auth_microservice/config"
	"auth_microservice/controllers"
	"auth_microservice/myMiddleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	config.LoadEnv()
	config.SmtpConfig()
	config.DbConfig()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.POST("/login", controllers.Login)
	e.POST("/sign-up", controllers.SignUp)
	e.GET("/logout", controllers.Logout)
	e.POST("/verify-email", controllers.VerifyEmail)
	t := e.Group("/token")
	t.Use(myMiddleware.JwtInterceptor)
	t.GET("", Hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello")
}
