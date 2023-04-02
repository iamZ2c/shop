package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	/*
		验证电话号码格式
	*/
	mobile := fl.Field().String()
	ok, err := regexp.MatchString("19823522875", mobile)
	zap.S().Info(ok)
	if err != nil {
		zap.S().Errorw("[validate err]", "msg", err.Error())
		return false
	}
	return ok
}
