package controller

import (
	// "net/http"
	"go/pkg/services"
	"github.com/gin-gonic/gin"
)

// ProductController : Product controller
type ProductController struct{}

func (c *ProductController) ProductList(c *gin.Context) {
	return services.ProductService
}
