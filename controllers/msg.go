package controllers

import (
	"encoding/json"
	"nim_163/models/msg"
	"nim_163/utils"
)

type MsgController struct {
	baseController
}

func (that *MsgController) Finish() {
	that.ServeJSON()
}

// 发送普通消息
// @router /sendMsg [post]
func (that *MsgController) SendMsg() {
	data := make(map[string]interface{})
	err := json.Unmarshal(that.Ctx.Input.RequestBody, &data)
	if err != nil {
		that.Data["json"] = utils.ResponseErr(err.Error())
		return
	}
	ope, _ := that.GetInt("ope", 0)
	sendType, _ := that.GetInt("sendType", 0)
	useYidun, _ := that.GetInt("useYidun")
	markRead, _ := that.GetInt("markRead")
	checkFriend, _ := that.GetBool("checkFriend")
	receiver := data["receiver"].([]interface{})
	content := data["content"]
	msgType := data["type"]
	attach := data["attache"]
	body := getMsgBody(content, msgType, attach)
	if len(receiver) == 1 {
		that.sendSingleMsg(receiver[0].(string), ope, sendType, useYidun, markRead, checkFriend, body)
	} else {
		// 一次最多100人
		for _, val := range utils.SliceChunk(receiver, 100) {
			receivers, err := json.Marshal(val)
			if err != nil {
				that.Data["json"] = utils.ResponseErr(err.Error())
				return
			}
			that.sendBatchMsg(string(receivers), sendType, useYidun, body)
		}
	}

}

// 生成json格式消息
func getMsgBody(content, msgType, attach interface{}) string {
	msgBody := make(map[string]interface{})
	body := make(map[string]interface{})
	body["type"] = msgType
	body["content"] = content
	body["attache"] = attach
	msgBody["msg"] = body
	jsonMsg, _ := json.Marshal(msgBody)
	return string(jsonMsg)
}

// 一对一发送消息
func (that *MsgController) sendSingleMsg(receiver string, ope, sendType, useYidun, markRead int, checkFriend bool, body string) {
	newMsg := msg.SendMsgStruct{
		From:             that.GetString("from", "admin"),
		Ope:              ope,
		To:               receiver,
		SendType:         sendType,
		Body:             body,
		Antispam:         that.GetString("antispam"),
		AntispamCustom:   that.GetString("antispamCustom"),
		Option:           that.GetString("option"),
		Pushcontent:      that.GetString("pushcontent"),
		Payload:          that.GetString("payload"),
		Ext:              that.GetString("ext"),
		Forcepushlist:    that.GetString("forcepushlist"),
		Forcepushcontent: that.GetString("forcepushcontent"),
		Forcepushall:     that.GetString("forcepushall"),
		Bid:              that.GetString("bid"),
		UseYidun:         useYidun,
		MarkRead:         markRead,
		CheckFriend:      checkFriend,
	}
	request := msg.SendMsg(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}

// 一对多发送消息
func (that *MsgController) sendBatchMsg(receiver string, sendType, useYidun int, body string) {
	returnMsgid, _ := that.GetBool("Boolean", true)
	newMsg := msg.SendBatchMsgStruct{
		FromAccid:   that.GetString("from", "admin"),
		ToAccids:    receiver,
		SendType:    sendType,
		Body:        body,
		Option:      that.GetString("option", msg.DefaultBatchMsgOption),
		Pushcontent: that.GetString("pushcontent"),
		Payload:     that.GetString("payload"),
		Ext:         that.GetString("ext"),
		Bid:         that.GetString("bid"),
		UseYidun:    useYidun,
		ReturnMsgid: returnMsgid,
	}
	request := msg.SendBatchMsg(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}

// 发送自定义系统通知
// @router /sendBatchAttachMsg [post]
func (that *MsgController) SendAttachMsg() {
	msgType, _ := that.GetInt("msgtype", 0)
	save, _ := that.GetInt("save", 2)
	newMsg := msg.SendAttachMsgStruct{
		From:        that.GetString("from", "admin"),
		MsgType:     msgType,
		To:          that.GetString("to"),
		Attach:      that.GetString("attach"),
		PushContent: that.GetString("pushcontent"),
		Payload:     that.GetString("payload"),
		Sound:       that.GetString("sound"),
		Save:        save,
		Option:      that.GetString("option", msg.DefaultAttachMsgOption),
	}
	request := msg.SendAttachMsg(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}

// 文件上传
// @router /upload [post]
func (that *MsgController) Upload() {
	expireSec, _ := that.GetInt64("expireSec")
	newMsg := msg.UploadStruct{
		Content:   that.GetString("content"),
		Type:      that.GetString("type"),
		IsHttps:   that.GetString("ishttps", "true"),
		ExpireSec: expireSec,
		Tag:       that.GetString("tag"),
	}
	request := msg.Upload(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}

// 消息撤回
// @router /recall [post]
func (that *MsgController) Recall() {
	timetag, _ := that.GetInt64("timetag")
	_type, _ := that.GetInt("type", 7)
	newMsg := msg.RecallStruct{
		DeleteMsgId: that.GetString("deleteMsgid"),
		TimeTag:     timetag,
		Type:        _type,
		From:        that.GetString("from"),
		To:          that.GetString("to"),
		Msg:         that.GetString("msg"),
		IgnoreTime:  that.GetString("ignoreTime", "1"),
		PushContent: that.GetString("pushcontent"),
		Payload:     that.GetString("payload"),
	}
	request := msg.Recall(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}

// 发送广播消息
// @router /broadcastMsg [post]
func (that *MsgController) BroadcastMsg() {
	ttl, _ := that.GetInt("ttl")
	newMsg := msg.BroadcastMsgStruct{
		Body:      that.GetString("body"),
		From:      that.GetString("from"),
		IsOffline: that.GetString("isOffline", "false"),
		Ttl:       ttl,
		TargetOs:  that.GetString("targetOs"),
	}
	request := msg.BroadcastMsg(newMsg)
	that.Data["json"] = utils.Request(request, newMsg)
}
