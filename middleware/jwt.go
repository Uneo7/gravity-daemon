package middleware

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetAuth(key *rsa.PublicKey) gin.HandlerFunc {

	return func(c *gin.Context) {
		header := c.Request.Header.Get("authorization")

		if header != "" {
			bearerToken := strings.Split(header, " ")

			if len(bearerToken) == 2 {

				token, _ := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
						return nil, fmt.Errorf("invalid token")
					}
					return key, nil
				})

				if token != nil && token.Valid {
					c.Next()
					return
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Access denied",
		})

		return
	}
}