package limiter

import (
	"errors"
	"net/http"
	"time"

	ajaxres "insulation/server/base/pkg/ajax_res"
	"insulation/server/base/pkg/config"
	"insulation/server/base/pkg/logger"
	redisutil "insulation/server/base/pkg/redis_util"
	"insulation/server/base/pkg/translater"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/text/message"
	"golang.org/x/time/rate"
)

var log *logger.Logger

func initLog() {
	log = logger.NewLogger("limiter", config.IsDebug())
}

func StandaloneLimiter(seconds int, tokens int) gin.HandlerFunc {
	if log == nil {
		initLog()
	}

	r := rate.NewLimiter(rate.Every(time.Duration(seconds)*time.Second), tokens)

	return func(c *gin.Context) {
		if !r.Allow() {
			tag, _ := translater.TranslaterFromContext(c)
			p := message.NewPrinter(tag)
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				ajaxres.Error(
					http.StatusTooManyRequests,
					p.Sprintf("请求过多"),
				),
			)

		} else {
			c.Next()
		}
	}
}

// redis 路由限流
func RedisLimiter(seconds int, tokens int) gin.HandlerFunc {
	if log == nil {
		initLog()
	}

	return func(ctx *gin.Context) {
		p := ctx.FullPath()
		var count int
		rk := redisutil.GenKey("limiter", "path", p)
		count, err := redisutil.GetInter[int](ctx.Request.Context(), rk)
		tag, _ := translater.TranslaterFromContext(ctx)
		tmsg := message.NewPrinter(tag)
		errMsg := tmsg.Sprintf("redis错误")
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				ajaxres.Error(http.StatusInternalServerError,
					errMsg,
				),
			)
		}
		if count > tokens {
			ctx.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				ajaxres.Error(
					http.StatusTooManyRequests,
					tmsg.Sprintf("请求过多"),
				),
			)
		} else {
			if count == 0 {
				err = redisutil.SetInterWithExprie(ctx.Request.Context(), rk, 1, time.Duration(seconds)*time.Second)
				if err != nil {
					log.Error(err.Error())
				}
			} else {
				err = redisutil.Incr(ctx.Request.Context(), rk)
				if err != nil {
					log.Error(err.Error())
				}
			}
			ctx.Next()
		}
	}
}

// redis ip限流
func RedisIpLimiter(seconds int, tokens int) gin.HandlerFunc {
	if log == nil {
		initLog()
	}

	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		p := ctx.FullPath()
		rk := redisutil.GenKey("limiter", "ip", ip, p)
		var count int
		count, err := redisutil.GetInter[int](ctx.Request.Context(), rk)
		tag, _ := translater.TranslaterFromContext(ctx)
		tmsg := message.NewPrinter(tag)
		errMsg := tmsg.Sprintf("redis错误")
		if err != nil && !errors.Is(err, redis.Nil) {
			log.Error(err.Error())
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				ajaxres.Error(http.StatusInternalServerError,
					errMsg,
				),
			)
		}
		if count > tokens {
			ctx.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				ajaxres.Error(
					http.StatusTooManyRequests,
					tmsg.Sprintf("请求过多"),
				),
			)
		} else {
			if count == 0 {
				err = redisutil.SetInterWithExprie(ctx.Request.Context(), rk, 1, time.Duration(seconds)*time.Second)
				if err != nil {
					log.Error(err.Error())
				}
			} else {
				err = redisutil.Incr(ctx.Request.Context(), rk)
				if err != nil {
					log.Error(err.Error())
				}
			}
			ctx.Next()
		}
	}
}
