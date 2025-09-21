package handlers

import (
	"net/http"
	"project/internal/services"

	"github.com/gin-gonic/gin"
)

type CsvHandler struct {
	services services.ICsvService
}

func NewCsvHandler(services services.ICsvService) *CsvHandler {
	return &CsvHandler{services: services}
}

func (h *CsvHandler) ProcessCsv(ctx *gin.Context) {
	err := h.services.ProcessCsv()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "CSV processed successfully"})
}
