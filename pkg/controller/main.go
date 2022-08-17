package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func resolve(c *gin.Context, req string, result interface{}, err map[string]interface{}) {
	if err != nil {
		go log.WithFields(log.Fields{
			"type":  "rest_api",
			"req":   req,
			"error": err,
		}).Error("Error")

		c.JSON(http.StatusOK, gin.H{"code": 400, "status": "error", "result": map[string]interface{}{}, "message": err["message"]})
		return
	}
	c.JSON(http.StatusOK, result)
}
