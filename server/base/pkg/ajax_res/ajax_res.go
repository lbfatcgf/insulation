package ajaxres

import (
	"insulation/server/base/pkg/translater"

	"golang.org/x/text/message"
)

// @Name			AjaxRes
//
// @Summary		ajax返回结构体
// @Description	ajax返回结构体
type AjaxRes struct {
	// 状态码
	Code int `json:"code" example:"200"`
	// 返回信息
	Msg string `json:"msg" example:"string"`
	// 返回数据
	Data any `json:"data"`
}

func Success(data any, lang string) AjaxRes {
	p := message.NewPrinter(translater.Translater(lang))

	success := p.Sprintf("成功")

	return AjaxRes{
		Code: 200,
		Msg:  success,
		Data: data,
	}
}

func Error(code int, msg string) AjaxRes {
	return AjaxRes{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
