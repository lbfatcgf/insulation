package routes

import (
	"net/http"

	"insulation/server/admin/internal/auth"
	ajaxres "insulation/server/base/pkg/ajax_res"
	"insulation/server/base/pkg/translater"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

func NewLoginRoute(engine *gin.Engine) {
	router := engine.Group("/auth")
	router.POST("/login", login())

	router.POST(`/text`, auth.JWTAuthMiddleware(), text())
}

// 后台用户登录
//
//	@Summary		后台用户登录
//	@Description	后台用户登录
//	@Tags			后台用户
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ajaxres.AjaxRes{data=string}
//	@Failure		403	{object}	ajaxres.AjaxRes{data=nil}
//	@Failure		500	{object}	ajaxres.AjaxRes{data=nil}
//	@Router			/auth/login [post]
func login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		langTag, lang := translater.TranslaterFromContext(ctx)
		p := message.NewPrinter(langTag)

		token, err := auth.GenerateToken(`{"name":"lbf"}`)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ajaxres.Error(http.StatusInternalServerError, p.Sprintf("凭证生成失败")))
			return
		}
		ctx.JSON(http.StatusOK, ajaxres.Success(token, lang))
	}
}

// 后台用户测试
//
//	@Summary		后台用户测试
//	@Description	后台用户测试
//	@Tags			后台用户
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	ajaxres.AjaxRes{data=string}
//	@Router			/auth/text [post]
func text() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, lang := translater.TranslaterFromContext(ctx)
		u := auth.GetAdminUser(ctx)
		ctx.JSON(http.StatusOK, ajaxres.Success(u, lang))
	}
}
