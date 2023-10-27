package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := IsTokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

const apiCode = "DASHORI"

func GenerateToken(login string, id uint64, role string) (string, error) {
	token_lifespan := 1

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["login"] = login
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(apiCode))
}

func IsTokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(apiCode), nil
	})

	return err
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenIdAndRole(c *gin.Context) (string, string, uint64, error) {
	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(apiCode), nil
	})

	if err != nil {
		return "", "", 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {

		login := fmt.Sprint(claims["login"])

		role := fmt.Sprint(claims["role"])

		id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)

		if err != nil {
			return "", "", 0, err
		}

		return login, role, id, nil
	}

	return "", "", 0, nil
}
