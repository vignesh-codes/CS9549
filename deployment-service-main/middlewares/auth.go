package middlewares

import (
	"deployment-service/apps/repository/adapter"
	"deployment-service/constants"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Application struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Service      string `json:"service"`
}

// for internal apis
func ValidateHeaderSecrets(repository *adapter.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// id := c.Request.Header.Get("x-client-id")
		// secret := c.Request.Header.Get("x-client-secret")
		// if id == "" {
		// 	status := response.UnAuthorized(string(response.ErrUnauthorized))
		// 	c.JSON(status.Status(), status)
		// 	c.Abort()
		// 	return
		// }
		// if secret == "" {
		// 	status := response.UnAuthorized(string(response.ErrUnauthorized))
		// 	c.JSON(status.Status(), status)
		// 	c.Abort()
		// 	return
		// }
		username := c.Request.Header.Get("username")
		c.Set("username", username)
		c.Next()
	}
}

func ValidateJWT(repository *adapter.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Bearer token parsing
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(constants.JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("err", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims and set the username in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if username, exists := claims["username"].(string); exists {
				c.Set("username", username)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found in token"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
