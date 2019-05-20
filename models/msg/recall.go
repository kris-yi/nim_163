package msg

import (
	"nim_163/models"
	"nim_163/server"
)

/**
消息撤回接口，可以撤回一定时间内的点对点与群消息
参数			类型	必须	说明
deleteMsgid		String	是	要撤回消息的msgid
timetag			long	是	要撤回消息的创建时间
type			int		是	7:表示点对点消息撤回，8:表示群消息撤回，其它为参数错误
from			String	是	发消息的accid
to				String	是	如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
msg				String	否	可以带上对应的描述
ignoreTime		String	否	1表示忽略撤回时间检测，其它为非法参数，如果需要撤回时间检测，不填即可
pushcontent		String	否	推送文案，android以此为推送显示文案；ios若未填写payload，显示文案以pushcontent为准。超过500字符后，会对文本进行截断。
payload			String	否	推送对应的payload,必须是JSON,不超过2K字符
*/
type RecallStruct struct {
	DeleteMsgId string `url:"deleteMsgid" valid:"Required"`
	TimeTag     int64  `url:"timetag" valid:"Required"`
	Type        int    `url:"type"`
	From        string `url:"from" valid:"Required"`
	To          string `url:"to" valid:"Required"`
	Msg         string `url:"msg"`
	IgnoreTime  string `url:"ignoreTime"`
	PushContent string `url:"pushcontent"`
	Payload     string `url:"payload"`
}

func Recall(recallStruct RecallStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "msg/recall.action",
			QueryData: recallStruct,
		},
	}
}
