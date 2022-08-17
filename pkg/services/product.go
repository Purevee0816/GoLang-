package services

import (
	"github.com/gin-gonic/gin"
	// "go/pkg/db"

	// "context"
)

type ProductService struct{}

func (u *ProductService) List(c *gin.Context) (r *Response) {

	// ctx := context.Background()
	// tx, _ := db.Tx(ctx)

	// var bank_list []models.AppBanks

	// var query = ""
	// if req.Name != "" {
	// 	query += "lower(name) = " + "'" + strings.ToLower(req.Name) + "'"
	// }

	// tx.Where(query).Order("id asc").Find(&bank_list)

	return r.SuccessWithParams(("success"))
}
