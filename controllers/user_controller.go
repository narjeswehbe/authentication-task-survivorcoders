package controllers

import (
	"auth_microservice/requests"
	"auth_microservice/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func SignUp(c echo.Context) error {
	var req requests.SighUpRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind request")
	}
	services.SignUp(req)
	return c.JSON(http.StatusCreated, "failed to create account")

}
func VerifyEmail(c echo.Context) error {
	var req requests.VerifyEmailRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind request")
	}
	services.VerifyEmail(req)
	return c.JSON(http.StatusCreated, "your account is verified")

}
func Login(c echo.Context) error {
	var req requests.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "failed to bind request")
	}
	token := services.Login(req)
	return c.JSON(http.StatusCreated, token)
}

func Logout(c echo.Context) error {
	headers := c.Request().Header
	auth := headers.Values("Authorization")
	if auth == nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	tokenString := strings.Fields(c.Request().Header.Get(echo.HeaderAuthorization))[1]
	ok := services.Logout(tokenString)
	if ok == false {
		return c.JSON(http.StatusBadRequest, "failed to log out")
	}
	return c.JSON(http.StatusOK, "logged out")

}
