package model

import (
	"WowjoyProject/WowjoyPdfArchive/pkg/setting"
	"database/sql"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type KeyData struct {
	ReportTime sql.NullString
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*sql.DB, error) {
	db, err := sql.Open(databaseSetting.DBType, databaseSetting.DBConn)
	if err != nil {
		return nil, err
	}
	// 数据库最大连接数
	db.SetConnMaxLifetime(time.Duration(databaseSetting.MaxLifetime) * time.Minute)
	db.SetMaxOpenConns(databaseSetting.MaxOpenConns)
	db.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	return db, nil
}
