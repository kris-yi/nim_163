package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
拉黑/取消拉黑；设置静音/取消静音
参数			类型	必须	说明
accid			String	是		用户帐号，最大长度32字符，必须保证一个APP内唯一
targetAcc		String	是		被加黑或加静音的帐号
relationType	int		是		本次操作的关系类型,1:黑名单操作，2:静音列表操作
value			int		是		操作值，0:取消黑名单或静音，1:加入黑名单或静音
*/
type SetSpecialRelationStruct struct {
	Accid        string `url:"accid" valid:"Required"`
	TargetAcc    string `url:"targetAcc" valid:"Required"`
	RelationType int    `url:"relationType" valid:"Required"`
	Value        int    `url:"value" valid:"Required"`
}

func SetSpecialRelation(setSpecialRelationStruct SetSpecialRelationStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/setSpecialRelation.action",
			QueryData: setSpecialRelationStruct,
		},
	}
}
