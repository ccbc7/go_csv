package middlewares

import (
	"net/http"
	"strings"

	"project/services"

	"github.com/gin-gonic/gin"
)

// HandlerFuncは、HTTPリクエストを処理するための関数を表す
func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Bearerトークンのチェック
		if !strings.HasPrefix(header, "Bearer ") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		// bearerを取り除いたトークンを取得し、ユーザーを取得
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)

		// ハンドラを続行
		ctx.Next()
	}
}
