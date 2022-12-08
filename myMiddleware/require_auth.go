package myMiddleware

import (
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

		if _, isValid := extractClaims(tokenString); isValid {
			return next(c)
		}
		return c.NoContent(http.StatusUnauthorized)

	}
}
func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
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
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
