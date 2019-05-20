package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数		类型	必须	说明
accid		String	是		用户帐号
donnopOpen	String	是		桌面端在线时，移动端是否不推送：true:移动端不需要推送，false:移动端需要推送
*/
type SetDesktopPushStruct struct {
	Accid      string `url:"accid" valid:"Required"`
	DonnopOpen string `url:"donnopOpen" valid:"Required"`
}

// 设置桌面端在线时，移动端是否需要推送
func SetDesktopPush(setDesktopPush SetDesktopPushStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/setDonnop.action",
			QueryData: setDesktopPush,
		},
	}
}
