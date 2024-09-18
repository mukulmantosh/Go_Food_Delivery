package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

type UserClaims struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

func ValidateToken(token string) (bool, int64) {

	tokenInfo, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		slog.Error("AuthMiddleware", "validateToken", err.Error())
		return false, 0
	} else if claims, ok := tokenInfo.Claims.(*UserClaims); ok {
		return true, claims.UserID
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return false, 0
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header!!"})
			c.Abort()
			return
		}

		// Extract Token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization"})
			c.Abort()
			return
		}
		token := tokenParts[1]

		tokenValidation, userID := ValidateToken(token)
		if !tokenValidation {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
