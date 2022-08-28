package util

import (
	"github.com/tidwall/gjson"
	"time"
)

const TIME_OFFSET_VALUE = -7

func AssignIntValue(jsonValue gjson.Result) (pValue *int) {
	if IsNull(jsonValue) || jsonValue.String() == "" {
		pValue = nil
	} else {
		value := int(jsonValue.Int())
		pValue = &value
	}

	return
}

func AssignInt64Value(jsonValue gjson.Result) (pValue *int64) {
	if IsNull(jsonValue) || jsonValue.String() == "" {
		pValue = nil
	} else {
		value := jsonValue.Int()
		pValue = &value
	}

	return
}

func AssignStringValue(jsonValue gjson.Result) (pValue *string) {
	if IsNull(jsonValue) {
		pValue = nil
	} else {
		value := jsonValue.String()
		pValue = &value
	}

	return
}

func AssignTimeValue(jsonValue gjson.Result) (pValue *time.Time) {
	if jsonValue.Int() == 0 {
		pValue = nil
	} else {
		timestampValue := jsonValue.Int()
		timestampValue = AddTimestamp(timestampValue, TIME_OFFSET_VALUE)
		value := time.Unix(timestampValue/1000, 0)
		pValue = &value
	}

	return
}

func AssignNotNilValue(value *int) (statusValue int) {
	statusValue = 0

	if value != nil {
		statusValue = *value
	}

	return
}
