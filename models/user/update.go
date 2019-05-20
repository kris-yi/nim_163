package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
更新用户名片
参数	类型	必须	说明
accid	String	是	用户帐号，最大长度32字符，必须保证一个APP内唯一
name	String	否	用户昵称，最大长度64字符，可设置为空字符串
icon	String	否	用户头像，最大长度1024字节，可设置为空字符串
sign	String	否	用户签名，最大长度256字符，可设置为空字符串
email	String	否	用户email，最大长度64字符，可设置为空字符串
birth	String	否	用户生日，最大长度16字符，可设置为空字符串
mobile	String	否	用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)，可设置为空字符串
gender	int		否	用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
ex		String	否	用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串，也可以设置为空字符串
*/
type UpdateStruct struct {
	Accid  string `url:"accid" valid:"Required"`
	Name   string `url:"name"`
	Icon   string `url:"icon"`
	Sign   string `url:"sign"`
	Email  string `url:"email"`
	Birth  string `url:"birth"`
	Mobile string `url:"mobile"`
	Gender string `url:"gender"`
	Ex     string `url:"ex"`
}

func Update(updateUser UpdateStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/updateUinfo.action",
			QueryData: updateUser,
		},
	}
}
