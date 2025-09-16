package controllers

import (
	"net/http"

	"project/internal/services"

	"project/internal/dto"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// サービス層のインターフェースを保持する構造体
type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &AuthController{service: service}
}

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the input payload
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.SignupInput	true	"User info"
//	@Success		201		{string}	string			"created"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		500		{string}	string			"internal server error"
//	@Router			/auth/signup [post]
func (c *AuthController) SignUp(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.SignUp(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// 成功時はステータスコード201を返却
	ctx.Status(http.StatusCreated)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login with the input payload
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.LoginInput	true	"User info"
//	@Success		200		{string}	string			"token"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		404		{string}	string			"user not found"
//	@Failure		500		{string}	string			"internal server error"
//	@Router			/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
