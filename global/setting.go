package global

import (
	"blog-service/pkg/logger"
	"blog-service/pkg/setting"
	ut "github.com/go-playground/universal-translator"
)

var (
	ServerSetting   *setting.ServerSettingS   //服务器配置
	AppSetting      *setting.AppSettingS      //App配置
	DatabaseSetting *setting.DatabaseSettingS //数据库配置

	Logger *logger.Logger // 日志
	Trans  ut.Translator  // 翻译器
)