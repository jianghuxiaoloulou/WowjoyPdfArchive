package v1

import (
	"WowjoyProject/WowjoyPdfArchive/global"
	"WowjoyProject/WowjoyPdfArchive/internal/model"
	"WowjoyProject/WowjoyPdfArchive/internal/routers/api/ws"
	"encoding/json"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 通过数据流上传文件
func ByFileStream(c *gin.Context) {
	// file, header, err := c.Request.FormFile("file")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "读取file失败",
		})
		return
	}
	global.Logger.Debug("接收到文件名：", file.Filename)
	//获取上传文件的后缀(类型)
	extType := path.Ext(file.Filename)
	// 文件名除去后缀 uid_enc
	uid_enc := strings.TrimSuffix(file.Filename, extType)
	// 增加判断是否是报告数据
	if !model.ExistData(uid_enc) {
		global.Logger.Debug("不是有效的文件名，", uid_enc)
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "不是报告数据,文件名不是有效UID_ENC",
		})
		return
	}
	global.Logger.Debug("是有效的文件名，", uid_enc)

	currentTime := time.Now()
	nYear := currentTime.Year()
	nMonth := currentTime.Month()
	nDay := currentTime.Day()

	timePath := "\\" + strconv.Itoa(nYear) + "\\" + strconv.Itoa(int(nMonth)) + "\\" + strconv.Itoa(nDay) + "\\"

	pdfData := global.PDFData{
		ReportId:      uid_enc,
		LocaltionCode: global.ObjectSetting.OBJECT_Upload_Success_Code,
		FileName:      timePath + file.Filename,
		ReportTime:    currentTime.Format("2006-01-02 15:04:05"),
	}

	// 文件保存路径
	saveDir := global.ObjectSetting.OBJECT_PDF_DIR + timePath
	// 打开目录
	localFileInfo, err := os.Stat(saveDir)
	// 目录不存在
	if err != nil || localFileInfo.IsDir() {
		// 创建目录
		err := os.MkdirAll(saveDir, 0755)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "创建目录失败",
			})
			data := global.Info{
				Code:   1,
				UidEnc: uid_enc,
				Msg:    "生成PDF失败",
			}
			info, _ := json.Marshal(data)
			ws.WebsocketManager.Send(uid_enc, "PDF", info)
		}
	}
	// 文件全路径
	savePath := saveDir + file.Filename
	global.Logger.Debug("文件保存路径：", saveDir)
	c.SaveUploadedFile(file, savePath)
	// 更新数据库
	model.UpdatePDF(pdfData)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "上传成功！",
	})
	data := global.Info{
		Code:   0,
		UidEnc: uid_enc,
		Msg:    "生成PDF成功",
	}
	info, _ := json.Marshal(data)
	ws.WebsocketManager.Send(uid_enc, "PDF", info)
}
