package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
//	@Summary	liveness probe
//	@Schemes
//	@Description	do ping
//	@Tags			Hello World
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	Helloworld
//	@Router			/hello [get]
func HelloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
