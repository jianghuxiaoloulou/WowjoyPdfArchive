package global

type PDFData struct {
	ReportId      string // 报告id
	LocaltionCode int    // 存储节点
	FileName      string // 文件名称
	ReportTime    string // 报告更新时间
}

type Info struct {
	Code   int    `json:"code"`
	UidEnc string `json:"uid_enc"`
	Msg    string `json:"message"`
}
