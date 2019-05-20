package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
1.自定义系统通知区别于普通消息，方便开发者进行业务逻辑的通知；
2.目前支持两种类型：点对点类型和群类型（仅限高级群），根据msgType有所区别。
应用场景：如某个用户给另一个用户发送好友请求信息等，具体attach为请求消息体，第三方可以自行扩展，建议是json格式
*/
type SendAttachMsgStruct struct {
	From        string `url:"from" valid:"Required"`
	MsgType     int    `url:"msgtype"`
	To          string `url:"to"`
	Attach      string `url:"attach"`
	PushContent string `url:"pushcontent"`
	Payload     string `url:"payload"`
	Sound       string `url:"sound"`
	Save        int    `url:"save"`
	Option      string `url:"option"`
}

const DefaultAttachMsgOption = "{'badge':false,'needPushNick':false,'route':false}"

func SendAttachMsg(msgStruct SendAttachMsgStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/sendAttachMsg.action",
			QueryData: msgStruct,
		},
	}
}
