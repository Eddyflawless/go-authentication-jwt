package routes

import (
	"net/http"

	ctr "go-jwt/api/controllers"
	mw "go-jwt/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())

	v1 := router.Group("/v1")

	AuthRoutes(v1)

	router.Use(mw.Jwt())

	UserRoutes(v1)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": "Access granted for api 1",
		})
	})

	router.GET("/me", ctr.Me)

	return router

}
