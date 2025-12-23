package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/VinceZCL/FinalYearProject/internal/service"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Missing Authorization header",
				})
			}

			// Split the header into "Bearer <token>"
			authParts := strings.Split(authHeader, " ")
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				// Invalid format, return Unauthorized
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Invalid Authorization header format. Expected 'Bearer <token>'",
				})
			}

			tokenStr := authParts[1]

			// Parse and validate the JWT token
			token, err := jwt.ParseWithClaims(tokenStr, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
				// Ensure the token's signing method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return service.JWTSecretKey, nil
			})

			// If parsing fails or the token is invalid, return Unauthorized
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Invalid or expired token",
				})
			}

			// Extract the claims
			claims, ok := token.Claims.(*service.Claims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Invalid token claims",
				})
			}

			// Optionally, check claims like expiration and issuer
			if claims.ExpiresAt.Time.Before(time.Now()) {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Token has expired",
				})
			}

			if claims.Issuer != "github.com/VinceZCL/FinalYearProject" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "Invalid issuer",
				})
			}

			c.Set("user", claims)

			return next(c)

		}
	}
}
