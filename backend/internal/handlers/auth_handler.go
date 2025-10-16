package handlers

import (
	"log"
	"net/http"

	"project/internal/services"

	"project/internal/dto"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

// サービス層のインターフェースを保持する構造体
type authHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) AuthHandler {
	return &authHandler{service: service}
}

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with the input payload
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.SignupInput	true	"Signup info"
//	@Success		201		{string}	string			"created"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		500		{string}	string			"internal server error"
//	@Router			/auth/signup [post]
func (h *authHandler) SignUp(ctx *gin.Context) {
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
//
//	@Summary		Login a user
//	@Description	Login a user with the input payload
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.LoginInput	true	"Login info"
//	@Success		200		{string}	string			"ok"
//	@Failure		400		{string}	string			"bad request"
//	@Failure		500		{string}	string			"internal server error"
//	@Router			/auth/login [post]
func (h *authHandler) Login(ctx *gin.Context) {
	log.Printf("=== LOGIN REQUEST START ===")

	var input dto.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		log.Printf("❌ JSON binding failed: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("📧 Received email: %s", input.Email)
	log.Printf("🔒 Password length: %d", len(input.Password))

	token, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		log.Printf("❌ Login service failed: %v", err)
		if err.Error() == "user not found" {
			log.Printf("👤 User not found for email: %s", input.Email)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Printf("💥 Internal error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("✅ Login successful for email: %s", input.Email)
	log.Printf("🎫 Token generated: %s...", (*token)[:20])
	log.Printf("=== LOGIN REQUEST END ===")

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
