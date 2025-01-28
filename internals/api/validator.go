package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sulabhghimire/bank-api/internals/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is supported or not'
		return util.IsSupportedCurrency(currency)
	}
	return false
}
