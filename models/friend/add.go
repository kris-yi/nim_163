package friend

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数		类型	必须	说明
accid		String	是	加好友发起者accid
faccid		String	是	加好友接收者accid
type		int		是	1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
msg			String	否	加好友对应的请求消息，第三方组装，最长256字符
serverex	String	否	服务器端扩展字段，限制长度256此字段client端只读，server端读写
*/
type AddStruct struct {
	Accid    string `url:"accid" valid:"Required"`
	Faccid   string `url:"faccid" valid:"Required"`
	Type     int    `url:"type" valid:"Required"`
	Msg      string `url:"msg"`
	Serverex string `url:"serverex"`
}

// 加好友
func Add(add AddStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "friend/add.action",
			QueryData: add,
		},
	}
}
