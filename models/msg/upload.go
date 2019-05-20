package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
文件上传，字符流需要base64编码，最大15M。
*/
type UploadStruct struct {
	Content   string `url:"content" valid:"Base64"`
	Type      string `url:"type"`
	IsHttps   string `url:"ishttps"`
	ExpireSec int64  `url:"expireSec valid:Mix(86400)"`
	Tag       string `url:"tag" valid:"MaxSize(32)"`
}

func Upload(uploadStruct UploadStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/upload.action",
			QueryData: uploadStruct,
		},
	}
}
