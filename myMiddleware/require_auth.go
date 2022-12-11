package myMiddleware

import (
	"auth_microservice/config"
	"auth_microservice/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"strings"
)

func JwtInterceptor(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		headers := c.Request().Header
		auth := headers.Values("Authorization")
		if auth == nil {
			return c.NoContent(http.StatusUnauthorized)
		}
		tokenString := strings.Fields(c.Request().Header.Get(echo.HeaderAuthorization))[1]

		if _, isValid := ExtractClaims(tokenString); isValid {
			return next(c)
		}
		return c.NoContent(http.StatusUnauthorized)

	}
}
func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := os.Getenv("SECRET_KEY")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var forbidden_token entity.BlackList
		// if this  token is blacklisted it is invalid , each means it did not expire yet but the user logged out
		config.Db.Where("auth_uuid = ?", claims["auth_uuid"]).First(&forbidden_token)
		if forbidden_token.ID != 0 {
			return claims, false
		}
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
