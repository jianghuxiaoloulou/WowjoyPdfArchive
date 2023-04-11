package setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GeneralSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt  string
	LogMaxSize  int
	LogMaxAge   int
}

type DatabaseSettingS struct {
	DBConn       string
	DBType       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  int
}

type ObjectSettingS struct {
	OBJECT_Upload_Success_Code int
	OBJECT_PDF_DIR             string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
