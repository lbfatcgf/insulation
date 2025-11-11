package translations

import "golang.org/x/text/message"

// 添加一个函数来标记需要翻译的字符串
func MarkMessages() {
	// 这些字符串将被gotext工具提取
	message.SetString(message.MatchLanguage("zh-CN"), "success", "成功")
	message.SetString(message.MatchLanguage("zh-CN"), "配证生成失败", "配证生成失败")
}

func init() {
	MarkMessages()
}
