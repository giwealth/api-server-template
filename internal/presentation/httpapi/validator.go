package httpapi

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// 解析自定义错误消息
func getError(err error, r interface{}) error {
	s := reflect.TypeOf(r)
	if s.Kind() == reflect.Ptr || s.Kind() == reflect.Interface {
		s = s.Elem()
	}

	var text string
	for _, fieldError := range err.(validator.ValidationErrors) {
		filed, _ := s.FieldByName(fieldError.Field())
		errText := filed.Tag.Get("error")
		if errText != "" {
			text += errText + ", "
		} else {
			text += fieldError.Field() + ":" + fieldError.Tag() + ", "
		}
	}

	return errors.New(strings.TrimSuffix(text, ", ") + ".")
}
