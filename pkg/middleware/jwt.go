package middleware

import (
	"context"
	"go-jwt/api/helpers"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractToken(r *http.Request) string {

	bearerToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Jwt() gin.HandlerFunc {

	return func(c *gin.Context) {

		// logic  here..
		accessToken := extractToken(c.Request)

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		}

		claims, err := helpers.DecodeJWT(accessToken)

		if err != nil {
			// somthing rather meaningful and with custom error code
			log.Printf("decoded JWT error: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		}

		userId := claims.Uid

		log.Printf("decoded access token %v\n", userId)

		if err != nil {
			// somthing rather meaningful and with custom error code
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "x_user_id", userId))

		c.Next()
	}
}
