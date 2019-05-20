package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数	类型	必须	说明
accid	String	是	用户帐号
mute	Boolean	是	是否全局禁言：true：全局禁言，false:取消全局禁言

*/
type MuteStruct struct {
	Accid string `url:"accid" valid:"Required"`
	Mute  bool   `url:"mute" valid:"Required"`
}

// 账号全局禁言
func Mute(mute MuteStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/mute.action",
			QueryData: mute,
		},
	}
}

// 账号全局禁用音视频
func MuteAv(mute MuteStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/muteAv.action",
			QueryData: mute,
		},
	}
}
