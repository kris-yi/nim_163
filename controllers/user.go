package controllers

import (
	"nim_163/models/user"
	"nim_163/utils"
)

type UserController struct {
	baseController
}

func (userController *UserController) Finish() {
	userController.ServeJSON()
}

// 创建网易云通信ID
// @router /create [post]
func (userController *UserController) Create() {
	newUser := user.CreateStruct{
		Accid:  userController.GetString("accid"),
		Name:   userController.GetString("name"),
		Props:  userController.GetString("props"),
		Icon:   userController.GetString("icon"),
		Token:  userController.GetString("token"),
		Sign:   userController.GetString("sign"),
		Email:  userController.GetString("email"),
		Birth:  userController.GetString("birth"),
		Mobile: userController.GetString("mobile"),
		Gender: userController.GetString("gender"),
		Ex:     userController.GetString("ex"),
	}
	request := user.Create(newUser)
	userController.Data["json"] = utils.Request(request, newUser)

}

// 更新用户名片
// @router /update [post]
func (userController *UserController) Update() {
	newUser := user.UpdateStruct{
		Accid:  userController.GetString("accid"),
		Name:   userController.GetString("name"),
		Icon:   userController.GetString("icon"),
		Sign:   userController.GetString("sign"),
		Email:  userController.GetString("email"),
		Birth:  userController.GetString("birth"),
		Mobile: userController.GetString("mobile"),
		Gender: userController.GetString("gender"),
		Ex:     userController.GetString("ex"),
	}
	request := user.Update(newUser)
	userController.Data["json"] = utils.Request(request, newUser)

}

// 设置桌面端在线时，移动端是否需要推送
// @router /setDonnop [post]
func (userController *UserController) SetDesktopPush() {
	newUser := user.SetDesktopPushStruct{
		Accid:      userController.GetString("accid"),
		DonnopOpen: userController.GetString("donnopOpen", "false"),
	}
	request := user.SetDesktopPush(newUser)
	userController.Data["json"] = utils.Request(request, newUser)
}

// 账号全局禁言
// @router /mute [post]
func (userController *UserController) Mute() {
	mute, err := userController.GetBool("mute")
	if err != nil {
		userController.Data["json"] = utils.ResponseErr("mute：" + err.Error())
		return
	}
	newUser := user.MuteStruct{
		Accid: userController.GetString("accid"),
		Mute:  mute,
	}
	request := user.Mute(newUser)
	userController.Data["json"] = utils.Request(request, newUser)
}

// 账号全局禁用音视频
// @router /muteAv [post]
func (userController *UserController) MuteAv() {
	mute, err := userController.GetBool("mute")
	if err != nil {
		userController.Data["json"] = utils.ResponseErr("mute：" + err.Error())
		return
	}
	newUser := user.MuteStruct{
		Accid: userController.GetString("accid"),
		Mute:  mute,
	}
	request := user.MuteAv(newUser)
	userController.Data["json"] = utils.Request(request, newUser)
}

//设置黑名单/静音
// @router /setSpecialRelation [post]
func (userController *UserController) SetSpacialRelation() {
	relationType, err := userController.GetInt("relationType")
	if err != nil {
		userController.Data["json"] = utils.ResponseErr(err.Error())
	}
	value, err := userController.GetInt("value")
	if err != nil {
		userController.Data["json"] = utils.ResponseErr(err.Error())
	}
	newUser := user.SetSpecialRelationStruct{
		Accid:        userController.GetString("accid"),
		TargetAcc:    userController.GetString("targetAcc"),
		RelationType: relationType,
		Value:        value,
	}
	request := user.SetSpecialRelation(newUser)
	userController.Data["json"] = utils.Request(request, newUser)
	userController.ServeJSON()
}

// 查看指定用户的黑名单和静音列表
// @router /listBlackAndMuteList [post]
func (userController *UserController) ListBlackAndMuteList() {
	newUser := user.ListBlackAndMuteListStruct{
		Accid: userController.GetString("accid"),
	}
	request := user.ListBlackAndMuteList(newUser)
	userController.Data["json"] = utils.Request(request, newUser)
}
