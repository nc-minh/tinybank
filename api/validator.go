package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nc-minh/tinybank/utils"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	return utils.IsSupportedCurrency(currency)
}
