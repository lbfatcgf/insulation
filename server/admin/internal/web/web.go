package web

import (
	"net/http"

	"insulation/server/admin/internal/auth"
	ajax_res "insulation/server/base/pkg/ajax_res"
	"insulation/server/base/pkg/translater"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

func Start(address string) {
	auth.InitializeAuth()
	app := gin.Default()

	app.POST("/login", func(ctx *gin.Context) {
		langTag, lang := translater.TranslaterFromContext(ctx)
		p := message.NewPrinter(langTag)

		token, err := auth.GenerateToken(`{"name":"lbf"}`)
		if err != nil {
			ctx.AbortWithStatusJSON(500, ajax_res.Error(http.StatusForbidden, p.Sprintf("凭证生成失败")))
			return
		}
		ctx.JSON(http.StatusOK, ajax_res.Success(token, lang))
	})

	app.POST(`/text`, auth.JWTAuthMiddleware(), func(ctx *gin.Context) {
		_, lang := translater.TranslaterFromContext(ctx)
		u := auth.GetAdminUser(ctx)
		ctx.JSON(http.StatusOK, ajax_res.Success(u, lang))
	})
	app.Run(address)
}
