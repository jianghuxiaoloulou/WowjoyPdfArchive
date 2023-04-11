package model

import (
	"WowjoyProject/WowjoyPdfArchive/global"
)

// 更新PDF_FILE表记录
func UpdatePDF(data global.PDFData) {
	global.Logger.Debug("开始更新PDF_FIlE ", data)
	sql := `INSERT INTO pdf_file(report_id,localtion_code,file_name) VALUE(?,?,?) ON DUPLICATE KEY UPDATE report_id = ?,localtion_code = ?,file_name = ?;`
	_, err := global.WriteDBEngine.Exec(sql, data.ReportId, data.LocaltionCode, data.FileName, data.ReportId, data.LocaltionCode, data.FileName)
	if err != nil {
		global.Logger.Error(err)
	}
}

// 判断是否是report 中的数据
func ExistData(uidenc string) bool {
	global.Logger.Debug("开始判断是否是report 表中记录", uidenc)
	sql := `select 1 from report where uid_enc = ? limit 1;`
	row := global.WriteDBEngine.QueryRow(sql, uidenc)
	err := row.Scan(&global.ExistStatus)
	if err != nil {
		global.Logger.Error(err)
		return false
	}
	if global.ExistStatus.Valid && global.ExistStatus.Int16 == 1 {
		return true
	}
	return false
}
