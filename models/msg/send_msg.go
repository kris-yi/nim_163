package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
给用户或者高级群发送普通消息，包括文本，图片，语音，视频和地理位置，具体消息参考下面链接地址
https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#发送普通消息
*/
type SendMsgStruct struct {
	From             string `url:"from" valid:"Required;MaxSize(32)"`
	Ope              int    `url:"ope"`
	To               string `url:"to" valid:"Required"`
	SendType         int    `url:"type"`
	Body             string `url:"body" valid:"Required;MaxSize(5000)"`
	Antispam         string `url:"antispam"`
	AntispamCustom   string `url:"antispamCustom"`
	Option           string `url:"option"`
	Pushcontent      string `url:"pushcontent"`
	Payload          string `url:"payload"`
	Ext              string `url:"ext"`
	Forcepushlist    string `url:"forcepushlist"`
	Forcepushcontent string `url:"forcepushcontent"`
	Forcepushall     string `url:"forcepushall"`
	Bid              string `url:"bid"`
	UseYidun         int    `url:"useYidun"`
	MarkRead         int    `url:"markRead"`
	CheckFriend      bool   `url:"checkFriend"`
}

func SendMsg(msgStruct SendMsgStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/sendMsg.action",
			QueryData: msgStruct,
		},
	}
}
