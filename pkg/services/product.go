package services

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

type ProductService struct{}

func (*ProductService) List(c *gin.Context) (r *Response) {
	c.JSON(http.StatusOK, "")
	c.Request.Context().Done()
}
