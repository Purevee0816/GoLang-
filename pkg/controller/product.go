package controller

import (
	"net/http"
	"go/pkg/services"
	"github.com/gin-gonic/gin"
)

// ProductController : Product controller
type ProductController struct{}

func (*ProductController) ProductList(c *gin.Context) {

	service := services.ProductService{}
	req, errS := c.GetRawData()
	reqS := string(req)
	if errS != nil {
		resolve(c, reqS, map[string]interface{}{}, map[string]interface{}{"error": errS, "message": "Алдаа"})
		return
	}
	result := service.List(reqS)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
	c.Request.Context().Done()
}
