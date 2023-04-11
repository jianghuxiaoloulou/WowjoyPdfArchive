package general

import (
	"os"
	"path/filepath"
)

// 基础函数
// action:对象操作类型：上传/下载/删除
// file :获取的对象名
// remotefile：数据库中保存的远端下载的key

type UploadFile struct {
	// 表单名称
	Name string
	// 文件全路径
	Filepath string
}

func GetFilePath(file, ip, virpath string) (path string) {
	path += "\\\\"
	path += ip
	path += "\\"
	path += virpath
	path += "\\"
	path += file
	return
}

func GetFileSize(filename string) int64 {
	fileInfo, err := os.Stat(filename)
	if err != nil {
	}
	return fileInfo.Size()
}

// 检查文件路径
func CheckPath(path string) {
	dir, _ := filepath.Split(path)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
