package routes

import (
	ctr "go-jwt/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {

	g1 := router.Group("/users")

	{
		g1.GET("/", ctr.GetUsers)
		g1.GET("/:userId", ctr.GetUser)
	}

}
