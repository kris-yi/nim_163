package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(response string, err error) ResponseMsg {
	if err != nil {
		return ResponseErr(err.Error())
	}
	var raw map[string]interface{}
	err = json.Unmarshal([]byte(response), &raw)
	if err != nil {
		return ResponseErr(err.Error())
	}
	if raw["code"] != float64(StatusOk) {
		return ResponseErr(fmt.Sprintf("%s", raw["desc"]))
	}
	beego.Info(raw)
	delete(raw, "code")
	return ResponseOk(raw)
}

func ResponseErr(desc string) ResponseMsg {
	return ResponseMsg{
		Code: StatusErr,
		Msg:  desc,
	}
}

func ResponseOk(desc interface{}) ResponseMsg {
	return ResponseMsg{
		Code: StatusOk,
		Msg:  "ok",
		Data: desc,
	}
}
