package middleware

import "github.com/gin-gonic/gin"

func Init(allowedRoles []string) gin.HandlerFunc {

	return func(c *gin.Context) {

		// todo: inject other dependencies here

		c.Next()
	}
}
