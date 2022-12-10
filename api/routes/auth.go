package routes

import (
	ctr "go-jwt/api/controllers"

	"github.com/gin-gonic/gin"
)

func signUpMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// logic  here..
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		// return
		c.Next()
	}
}

func AuthRoutes(router *gin.RouterGroup) {

	g1 := router.Group("/auth")

	g1.Use(signUpMiddleware())
	{

		g1.POST("signup", ctr.Signup)
		g1.POST("login", ctr.Login)
	}

}
