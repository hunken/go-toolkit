package util

import "reflect"

// IsNil check if interface is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice, reflect.Interface, reflect.Func, reflect.UnsafePointer:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// IsStringSliceContains check if `arr` string slice contains `find` string or not
func IsStringSliceContains(arr []string, find string) bool {
	for _, s := range arr {
		if find == s {
			return true
		}
	}

	return false
}

// CensorString return censored string with some distract char
func CensorString(secret string) string {
	dlen := len(secret) / 2
	if dlen > maxClearTextLen {
		dlen = maxClearTextLen
	}
	dchars := make([]byte, dlen)
	for i := 0; i < dlen; i++ {
		dchars[i] = secret[i]
	}
	dchars = append(dchars, '*', '*', '*')

	return string(dchars)
}

// CensorPhone return censored phone by X characters except first `prefix` chars and last `suffix` end chars
func CensorPhone(phone string, prefix int, suffix int) string {
	dchars := make([]byte, len(phone))

	for i := 0; i < len(phone); i++ {
		if i < prefix {
			dchars[i] = phone[i]
		} else if i >= len(phone)-suffix {
			dchars[i] = phone[i]
		} else {
			dchars[i] = 'X'
		}
	}

	return string(dchars)
}
