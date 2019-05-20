package user

import (
	"nim_163/models"
	"nim_163/server"
)

/**
参数	类型	必须	说明
accid	String	是	网易云通信ID，最大长度32字符，必须保证一个APP内唯一（只允许字母、数字、半角下划线_、@、半角点以及半角-组成，不区分大小写，会统一小写处理，请注意以此接口返回结果中的accid为准）。
name	String	否	网易云通信ID昵称，最大长度64字符，用来PUSH推送时显示的昵称
props	String	否	json属性，第三方可选填，最大长度1024字符
icon	String	否	网易云通信ID头像URL，第三方可选填，最大长度1024
token	String	否	网易云通信ID可以指定登录token值，最大长度128字符，并更新，如果未指定，会自动生成token，并在创建成功后返回
sign	String	否	用户签名，最大长度256字符
email	String	否	用户email，最大长度64字符
birth	String	否	用户生日，最大长度16字符
mobile	String	否	用户mobile，最大长度32字符，非中国大陆手机号码需要填写国家代码(如美国：+1-xxxxxxxxxx)或地区代码(如香港：+852-xxxxxxxx)
gender	int		否	用户性别，0表示未知，1表示男，2女表示女，其它会报参数错误
ex		String	否	用户名片扩展字段，最大长度1024字符，用户可自行扩展，建议封装成JSON字符串
*/
type CreateStruct struct {
	Accid  string `url:"accid" valid:"Required"`
	Name   string `url:"name"`
	Props  string `url:"props"`
	Icon   string `url:"icon"`
	Token  string `url:"token"`
	Sign   string `url:"sign"`
	Email  string `url:"email"`
	Birth  string `url:"birth"`
	Mobile string `url:"mobile"`
	Gender string `url:"gender"`
	Ex     string `url:"ex"`
}

// 创建网易云通信ID
func Create(createUser CreateStruct) *models.BaseRequest {
	return &models.BaseRequest{
		NimRequest: &server.NimRequest{
			Api:       "user/create.action",
			QueryData: createUser,
		},
	}
}
