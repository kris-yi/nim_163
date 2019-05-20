package friend

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数		类型	必须	说明
accid		String	是		发起者accid
faccid		String	是		要修改朋友的accid
alias		String	否		给好友增加备注名，限制长度128，可设置为空字符串
ex			String	否		修改ex字段，限制长度256，可设置为空字符串
serverex	String	否		修改serverex字段，限制长度256，可设置为空字符串此字段client端只读，server端读写
*/
type UpdateStruct struct {
	Accid    string `url:"accid"`
	Faccid   string `url:"faccid"`
	Alias    string `url:"alias"`
	Ex       string `url:"ex"`
	Serverex string `url:"serverex"`
}

// 更新好友相关信息
func Update(update UpdateStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "friend/update.action",
			QueryData: update,
		},
	}
}
