package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
查看用户的黑名单和静音列表
参数	类型	必须	说明
accid	String	是		用户帐号，最大长度32字符，必须保证一个APP内唯一
*/

type ListBlackAndMuteListStruct struct {
	Accid string `url:"accid" valid:"Required;MaxSize(32)"`
}

func ListBlackAndMuteList(listStruct ListBlackAndMuteListStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/listBlackAndMuteList.action",
			QueryData: listStruct,
		},
	}
}
