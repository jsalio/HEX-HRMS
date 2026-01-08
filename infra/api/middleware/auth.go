package middleware

import (
	"fmt"
	"hrms/core/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	secretKey []byte
	Config    *AuthConfig
}

func NewAuthMiddleware() *AuthMiddleware {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "3xtr3m4d4m3nt3C0mpl3j0"
	}
	return &AuthMiddleware{
		secretKey: []byte(secretKey),
		Config:    NewAuthConfig(),
	}
}

func (am *AuthMiddleware) SkipAuth(c *gin.Context, skipRoutes []string) bool {
	path := c.FullPath()
	method := c.Request.Method
	if method == "OPTIONS" {
		return true
	}
	for _, route := range skipRoutes {
		if route == path {
			if strings.HasPrefix(path, route) && (method == "POST" || route == "/swagger/") {
				return true
			}
		}
	}
	return false
}

func (m *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var authError = models.SystemError{
			Code:    http.StatusUnauthorized,
			Message: "Se requiere token de autenticación",
		}

		if m.SkipAuth(c, m.Config.PublicRoutes) {
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, authError)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return m.secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["sub"])
			c.Set("data", claims["data"])
		}

		c.Next()
	}
}

func (m *AuthMiddleware) GenerateToken(userID string, data map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"data": data,
		"exp":  jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})
	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
