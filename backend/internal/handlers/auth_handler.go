package handlers

import (
	"net/http"

	"project/internal/services"

	"project/internal/dto"

	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// サービス層のインターフェースを保持する構造体
type AuthHandler struct {
	service services.IAuthService
}

func NewAuthHandler(service services.IAuthService) IAuthHandler {
	return &AuthHandler{service: service}
}

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the input payload
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.SignupInput	true	"User info"	example({"name": "John Doe", "email": "john@example.com", "password": "password123"})
//	@Success		201		{string}	string			"created"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		500		{string}	string			"internal server error"
//	@Router			/auth/signup [post]
func (h *AuthHandler) SignUp(ctx *gin.Context) {
	var input dto.SignupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.SignUp(input.Name, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// 成功時はステータスコード201を返却
	ctx.Status(http.StatusCreated)
}

// Login godoc
func (h *AuthHandler) Login(ctx *gin.Context) {
	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(input.Email, input.Password)
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
