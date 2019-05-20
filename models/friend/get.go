package friend

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数		类型	必须	说明
accid		String	是		发起者accid
updatetime	Long	是		更新时间戳，接口返回该时间戳之后有更新的好友列表
createtime	Long	否		【Deprecated】定义同updatetime
*/
type GetStruct struct {
	Accid      string `url:"accid" valid:"Required"`
	UpdateTime int64  `url:"updatetime" valid:"Required"`
	CreateTime int64  `url:"createtime"`
}

// 获取好友关系
func Get(get GetStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "friend/get.action",
			QueryData: get,
		},
	}
}
