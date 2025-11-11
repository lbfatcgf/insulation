package translations

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

var langSupport = []string{`zh_CN`, `en_US`}

func Translater(lang string) language.Tag {
	key := strings.ReplaceAll(lang, "-", "_")
	switch key {
	case `zh_CN`:
		return language.MustParse(`zh-CN`)
	case `en_US`:
		return language.MustParse(`en-US`)
	default:
		return language.MustParse(`zh-CN`)
	}
}

// 从http头Accept-Language或get参数lang获取地区码
func TranslaterFromContext(ctx *gin.Context) (language.Tag, string) {
	lang := ctx.GetHeader(`Accept-Language`)
	if lang == "" {

		lang = ctx.Query(`lang`)
		if lang == "" {
			lang = `zh_CN`
		}
	}
	lang = strings.ReplaceAll(lang, "-", `_`)
	if !slices.Contains(langSupport, lang) {
		lang = `zh_CN`
	}
	return Translater(lang), lang
}
