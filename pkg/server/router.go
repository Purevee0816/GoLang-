package server

import (
	"go/pkg/controller"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "okey")
	})

	productGroup := router.Group("/product")
	{
		var product controller.ProductController
		productGroup.GET("/:id", product.ProductList)
	}

	// var dan controller.DanController
	// router.GET("/dan", dan.Make)

	return router
}
