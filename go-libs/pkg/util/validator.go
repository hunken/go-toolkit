package util

import "github.com/tidwall/gjson"

func IsNull(value gjson.Result) (result bool) {
	result = false

	if value.Value() == nil {
		result = true
	}

	return
}
