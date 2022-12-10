package middleware

import "github.com/gin-gonic/gin"

func Authorize(allowedRoles []string) gin.HandlerFunc {

	return func(c *gin.Context) {

		// logic here..

		c.Next()
	}
}
