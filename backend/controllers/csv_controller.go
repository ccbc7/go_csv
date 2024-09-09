package controllers

import (
	"net/http"
	"project/services"

	"github.com/gin-gonic/gin"
)

type CsvController struct {
	services services.ICsvService
}

func NewCsvController(services services.ICsvService) *CsvController {
	return &CsvController{services: services}
}

func (c *CsvController) ProcessCsv(ctx *gin.Context) {
	err := c.services.ProcessCsv()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "CSV processed successfully"})
}

