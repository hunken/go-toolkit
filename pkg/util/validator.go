package util

import (
	"github.com/asaskevich/govalidator"
	"github.com/tidwall/gjson"
	"net"
	"net/url"
	"regexp"
	"strings"
)

func IsNull(value gjson.Result) (result bool) {
	result = false

	if value.Value() == nil {
		result = true
	}

	return
}

const (
	MobilePhone = "^84(3|5|7|8|9)([0-9]{8})$"
	Phone       = "^84((3|5|7|8|9)([0-9]{8})|(2[0-9]{9}))$"
	Username    = "^[a-zA-Z0-9][-.a-zA-Z0-9]+$"
)

var (
	rxMobilePhone = regexp.MustCompile(MobilePhone)
	rxPhone       = regexp.MustCompile(Phone)
	rxUsername    = regexp.MustCompile(Username)
)

// IsPhone check if a string is a valid Vietnamese phone number or not
func IsPhone(s string) bool {
	return !IsStringEmpty(s) && rxPhone.MatchString(s)
}

// IsMobilePhone check if a string is a valid Vietnamese mobile phone number or not
func IsMobilePhone(s string) bool {
	return !IsStringEmpty(s) && rxMobilePhone.MatchString(s)
}

func IsEmail(s string) bool {
	return !IsStringEmpty(s) && govalidator.IsEmail(s)
}

func IsValidUsername(s string) bool {
	return !IsStringEmpty(s) && rxUsername.MatchString(s)
}

func ValidateURL(urlStr string) bool {
	return ValidateRawURL(urlStr, true)
}

func ValidateRawURL(urlStr string, requireSSL bool) bool {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	if requireSSL {
		if parsedURL.Scheme != "https" {
			return false
		}
	}
	if IsStringEmpty(parsedURL.Host) {
		return false
	}
	s := strings.Split(parsedURL.Host, ".")
	if IsStringEmpty(s[0]) {
		return false
	}
	if requireSSL {
		if len(s) <= 1 {
			return false
		}
		host := s[len(s)-1]
		if strings.Contains(host, ":") {
			host, _, err = net.SplitHostPort(host)
			if err != nil {
				return false
			}
		}
		if host == "localhost" {
			return false
		}
	}
	return govalidator.IsURL(urlStr)
}
