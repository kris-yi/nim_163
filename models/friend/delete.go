package friend

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数			类型	必须	说明
accid			String	是		发起者accid
faccid			String	是		要删除朋友的accid
isDeleteAlias	Boolean	否		是否需要删除备注信息默认false:不需要，true:需要
*/
type DeleteStruct struct {
	Accid         string `url:"accid"`
	Faccid        string `url:"faccid"`
	IsDeleteAlias bool   `url:"isDeleteAlias"`
}

// 删除好友
func Delete(delete DeleteStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "friend/delete.action",
			QueryData: delete,
		},
	}
}
