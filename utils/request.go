package utils

import (
	"nim_163/server"
)

func Request(request server.Request, val interface{}) ResponseMsg {
	str := Validate(val)
	if str != "" {
		return ResponseErr(str)
	}
	client, _ := server.ClientInstance()
	response, err := client.Do(request)
	return Response(response, err)
}
