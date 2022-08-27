package util

import (
	"regexp"
	"strings"
)

var TEL_PREFIX_CONVERT = map[string]string{
	"120": "70", "121": "79", "122": "77", "126": "76", "128": "78",
	"123": "83", "124": "84", "125": "85", "127": "81", "129": "82",
	"162": "32", "163": "33", "164": "34", "165": "35", "166": "36",
	"167": "37", "168": "38", "169": "39", "186": "56", "188": "58",
	"199": "59",
}

func ConvertPhoneNumber(originalTel *string) *string {
	if originalTel == nil {
		return nil
	}
	normalizeNumber := RemoveNonDigits(*originalTel)
	normalizeNumber = RemoveLeadingZero(normalizeNumber)
	if len(normalizeNumber) >= 12 {
		reStr := regexp.MustCompile("^(840|84)*(.*)")
		repl := "$2"
		normalizeNumber = reStr.ReplaceAllString(normalizeNumber, repl)
	}
	normalizeNumber = RemoveLeadingZero(normalizeNumber)
	if len(normalizeNumber) < 8 || len(normalizeNumber) > 11 {
		return &normalizeNumber
	}
	prefix := normalizeNumber[0:3]
	convertedPrefix := ""
	if val, found := TEL_PREFIX_CONVERT[prefix]; found {
		convertedPrefix = val
	}
	if convertedPrefix != "" {
		normalizeNumber = convertedPrefix + normalizeNumber[3:]
	}
	numberLength := len(normalizeNumber)
	IsCallCenter := numberLength == 8 && (strings.HasPrefix(normalizeNumber, "18") || strings.HasPrefix(normalizeNumber, "19"))
	IsLandline := numberLength == 10 && strings.HasPrefix(normalizeNumber, "2")
	IsMobileNumber := numberLength == 9
	if IsCallCenter || IsLandline || IsMobileNumber {
		normalizeNumber = "84" + normalizeNumber
		return &normalizeNumber
	}

	return &normalizeNumber
}

func RemoveNonDigits(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	return strings.ReplaceAll(result.String(), " ", "")
}

func RemoveLeadingZero(s string) string {
	return strings.TrimLeft(s, "0")
}
