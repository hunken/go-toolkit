package util

import (
	"net"
	"net/url"
	"strings"
)

func CheckURLWhitelist(urlStr string, whiteListDomains []string) bool {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return false
	}
	host := parsedURL.Host

	for _, whitelistDomain := range whiteListDomains {
		// Allow wildcard such as *.xxx.vn
		if strings.Contains(whitelistDomain, "*") &&
			strings.HasSuffix(host, strings.Replace(whitelistDomain, "*", "", 1)) {
			return true
		} else if host == whitelistDomain {
			return true
		}
	}

	return false
}

func CheckBlacklistDomain(urlStr string, blackListDomains []string) bool {
	parsedURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return true
	}
	host := parsedURL.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return true
		}
	}
	if util.IsStringSliceContains(blackListDomains, host) {
		return true
	}
	for _, bd := range blackListDomains {
		if strings.HasSuffix(host, "."+bd) {
			return true
		}
	}
	return false
}
