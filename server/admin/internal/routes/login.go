package routes

import (
	"net/http"

	"insulation/server/admin/internal/auth"
	ajax_res "insulation/server/base/pkg/ajax_res"
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/limiter"
	"insulation/server/base/pkg/logger"
	"insulation/server/base/pkg/translater"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

var log *logger.Logger

func NewLoginRoute(engine *gin.Engine) {
	log = logger.NewLogger("login", config.IsDebug())

	router := engine.Group("/auth")
	router.Use(limiter.RedisIpLimiter(1, 100))
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
		log.Info(lang)
		token, err := auth.GenerateToken(`{"name":"lbf"}`)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ajax_res.Error(http.StatusInternalServerError, p.Sprintf("凭证生成失败")))
			return
		}
		ctx.JSON(http.StatusOK, ajax_res.Success(token, lang))
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
		log.Info(ctx.Request.RequestURI)
		_, lang := translater.TranslaterFromContext(ctx)
		u := auth.GetAdminUser(ctx)
		ctx.JSON(http.StatusOK, ajax_res.Success(u, lang))
	}
}
