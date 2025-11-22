package auth

import (
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/jwt_util"

	"github.com/gin-gonic/gin"
)

var jwt *jwt_util.AdminJwt

func InitializeAuth() {
	opt := config.Global().JwtOptions
	j := jwt_util.NewAdminJwt()
	err := j.SetSecret(opt)
	if err != nil {
		panic(err)
	}
	jwt = j
}

func GenerateToken(payload string) (string, error) {
	return jwt.GenerateToken(payload)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "token is empty",
			})
			return
		}
		payload, err := jwt.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"message": "token is invalid",
			})
			return
		}
		SetAdminUser(ctx, payload)
		ctx.Next()
	}
}

func SetAdminUser(ctx *gin.Context, user string) {
	ctx.Set("admin_user", user)
}

func GetAdminUser(ctx *gin.Context) string {
	return ctx.GetString("admin_user")
}
