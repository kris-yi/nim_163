package utils

import (
	"github.com/astaxie/beego/validation"
	"strings"
)

func Validate(val interface{}) string {
	valid := validation.Validation{}
	v, err := valid.Valid(val)
	var str []string
	if err != nil {
		return err.Error()
	}
	if !v {
		for _, err := range valid.Errors {
			str = append(str, err.Key+","+err.Message)
		}
	}
	return strings.Join(str, ";")
}
