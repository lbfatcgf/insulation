package web

import (
	"net/http"

	"insulation/server/admin/internal/auth"
	ajax_res "insulation/server/base/pkg/ajax_res"
	"insulation/server/base/pkg/translations"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

func Start(address string) {
	auth.InitializeAuth()
	app := gin.Default()

	app.POST("/login", func(ctx *gin.Context) {
		langTag, lang := translations.TranslaterFromContext(ctx)
		translater := message.NewPrinter(langTag)
		token, err := auth.GenerateToken(`{"name":"lbf"}`)
		if err != nil {
			ctx.AbortWithStatusJSON(500, ajax_res.Error(http.StatusForbidden, translater.Sprintf("配证生成失败")))
			return
		}
		ctx.JSON(http.StatusOK, ajax_res.Success(token, lang))
	})

	app.POST(`/text`, auth.JWTAuthMiddleware(), func(ctx *gin.Context) {
		_, lang := translations.TranslaterFromContext(ctx)
		u := auth.GetAdminUser(ctx)
		ctx.JSON(http.StatusOK, ajax_res.Success(u, lang))
	})
	app.Run(address)
}
