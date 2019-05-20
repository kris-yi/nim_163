package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
1、广播消息，可以对应用内的所有用户发送广播消息，广播消息目前暂不支持第三方推送（APNS、小米、华为等）；
2、广播消息支持离线存储，并可以自定义设置离线存储的有效期，最多保留最近100条离线广播消息；
3、此接口受频率控制，一个应用一分钟最多调用10次，一天最多调用1000次，超过会返回416状态码；
4、该功能目前需申请开通，详情可咨询您的客户经理。

参数		类型	必须	说明
body		String	是	广播消息内容，最大4096字符
from		String	否	发送者accid, 用户帐号，最大长度32字符，必须保证一个APP内唯一
isOffline	String	否	是否存离线，true或false，默认false
ttl			int		否	存离线状态下的有效期，单位小时，默认7天
targetOs	String	否	目标客户端，默认所有客户端，jsonArray，格式：["ios","aos","pc","web","mac"]
*/
type BroadcastMsgStruct struct {
	Body      string `url:"body" valid:"Required;MaxSize(4096)"`
	From      string `url:"from"`
	IsOffline string `url:"isOffline"`
	Ttl       int    `url:"ttl"`
	TargetOs  string `url:"targetOs"`
}

func BroadcastMsg(msgStruct BroadcastMsgStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/broadcastMsg.action",
			QueryData: msgStruct,
		},
	}
}
