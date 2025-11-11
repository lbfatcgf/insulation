package ajaxres

import (
	"insulation/server/base/pkg/translater"

	"golang.org/x/text/message"
)

type AjaxRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
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
