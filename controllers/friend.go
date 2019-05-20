package controllers

import (
	"nim_163/models/friend"
	"nim_163/utils"
	"strconv"
	"strings"
)

type FriendController struct {
	baseController
}

func (friendController *FriendController) Finish() {
	friendController.ServeJSON()
}

// 加好友
// @router /add [post]
func (friendController *FriendController) Add() {
	friendType, err := strconv.Atoi(strings.TrimSpace(friendController.GetString("type")))
	if err != nil {
		friendController.Data["json"] = utils.ResponseErr("type：" + err.Error())
		return
	}
	newFriend := friend.AddStruct{
		Accid:    friendController.GetString("accid"),
		Faccid:   friendController.GetString("faccid"),
		Type:     friendType,
		Msg:      friendController.GetString("msg"),
		Serverex: friendController.GetString("serverex"),
	}
	request := friend.Add(newFriend)
	friendController.Data["json"] = utils.Request(request, newFriend)
}

//更新好友相关信息
// @router /update [post]
func (friendController *FriendController) Update() {
	accid := friendController.GetString("accid")
	faccid := friendController.GetString("faccid")
	if accid == "" || faccid == "" {
		friendController.Data["json"] = utils.ResponseErr("accid或faccid不能为空")
		return
	}
	newFriend := friend.UpdateStruct{
		Accid:    accid,
		Faccid:   faccid,
		Alias:    friendController.GetString("alias"),
		Ex:       friendController.GetString("ex"),
		Serverex: friendController.GetString("serverex"),
	}
	request := friend.Update(newFriend)
	friendController.Data["json"] = utils.Request(request, newFriend)
}

//删除好友
// @router /delete [post]
func (friendController *FriendController) Delete() {
	accid := friendController.GetString("accid")
	faccid := friendController.GetString("faccid")
	isDeleteAlias, err := strconv.ParseBool(strings.TrimSpace(friendController.GetString("isDeleteAlias", "true")))
	if accid == "" || faccid == "" {
		friendController.Data["json"] = utils.ResponseErr("accid或faccid不能为空")
		return
	}
	if err != nil {
		friendController.Data["json"] = utils.ResponseErr("isDeleteAlias：" + err.Error())
		return
	}
	newFriend := friend.DeleteStruct{
		Accid:         accid,
		Faccid:        faccid,
		IsDeleteAlias: isDeleteAlias,
	}
	request := friend.Delete(newFriend)
	friendController.Data["json"] = utils.Request(request, newFriend)
}

//获取好友关系
// @router /get [post]
func (friendController *FriendController) Get() {
	accid := friendController.GetString("accid")
	updatetime, _ := friendController.GetInt64("updatetime")
	createtime, _ := friendController.GetInt64("createtime")
	newFriend := friend.GetStruct{
		Accid:      accid,
		UpdateTime: updatetime,
		CreateTime: createtime,
	}
	request := friend.Get(newFriend)
	friendController.Data["json"] = utils.Request(request, newFriend)
}
