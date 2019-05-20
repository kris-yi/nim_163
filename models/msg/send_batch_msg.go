package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
1.给用户发送点对点普通消息，包括文本，图片，语音，视频，地理位置和自定义消息。
2.最大限500人，只能针对个人,如果批量提供的帐号中有未注册的帐号，会提示并返回给用户。
3.此接口受频率控制，一个应用一分钟最多调用120次，超过会返回416状态码，并且被屏蔽一段时间；具体消息参考下面描述。
https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#发送普通消息
*/
type SendBatchMsgStruct struct {
	FromAccid   string `url:"fromAccid" valid:"Required;MaxSize(32)"`
	ToAccids    string `url:"toAccids" valid:"Required"`
	SendType    int    `url:"type"`
	Body        string `url:"body" valid:"Required"`
	Option      string `url:"option"`
	Pushcontent string `url:"pushcontent"`
	Payload     string `url:"payload"`
	Ext         string `url:"ext" valid:"MaxSize(1024)"`
	Bid         string `url:"bid"`
	UseYidun    int    `url:"useYidun"`
	ReturnMsgid bool   `url:"returnMsgid"`
}

const DefaultBatchMsgOption = "{'push':false,'roam':true,'history':false,'sendersync':true,'route':false,'badge':false,'needPushNick':true}"

func SendBatchMsg(msgStruct SendBatchMsgStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/sendBatchMsg.action",
			QueryData: msgStruct,
		},
	}
}
