package global

import (
	"WowjoyProject/WowjoyPdfArchive/pkg/logger"
	"WowjoyProject/WowjoyPdfArchive/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	GeneralSetting  *setting.GeneralSettingS
	DatabaseSetting *setting.DatabaseSettingS
	ObjectSetting   *setting.ObjectSettingS
	Logger          *logger.Logger
)
