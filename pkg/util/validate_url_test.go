package utilpackage util

import (
"reflect"
"testing"
)

type testCaseURL struct {
	name    string
	testURL string
	want    bool
}

func TestCheckURLWhitelist(t *testing.T) {
	whiteList := []string{"ghtk.vn", "ghtklab.com", "cdn.image.ghtk.vn", "*.wildcard.vn"}
	tests := []testCaseURL{
		{
			name:    "fail test case 1",
			testURL: "https://google.com/image.png",
			want:    false,
		},
		{
			name:    "fail test case 2",
			testURL: "https://ghtk.vn.google.com/image.png",
			want:    false,
		},
		{
			name:    "fail test case 3",
			testURL: "https://ghtk.vn.google.com/image.png",
			want:    false,
		},
		{
			name:    "success test case 1",
			testURL: "https://ghtk.vn/image.png",
			want:    true,
		},
		{
			name:    "success test case 2",
			testURL: "https://cdn.image.ghtk.vn/image.png",
			want:    true,
		},
		{
			name:    "test wildcard success",
			testURL: "https://hihi.wildcard.vn/image.png",
			want:    true,
		},
		{
			name:    "test wildcard fail",
			testURL: "https://hihi.dwildcard.vn/image.png",
			want:    false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(subT *testing.T) {
			got := CheckURLWhitelist(test.testURL, whiteList)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: expect mismatch, got: %v want: %v", test.name, got, test.want)
			}
		})
	}
}

func TestCheckBlacklistDomain(t *testing.T) {
	type args struct {
		urlStr           string
		blackListDomains []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not match",
			args: args{
				urlStr:           "https://apple.com",
				blackListDomains: []string{"google.com", "abc.com"},
			},
			want: false,
		},
		{
			name: "match exactly",
			args: args{
				urlStr:           "https://google.com",
				blackListDomains: []string{"google.com", "abc.com"},
			},
			want: true,
		},
		{
			name: "match exactly",
			args: args{
				urlStr:           "http://localhost:8080",
				blackListDomains: []string{"google.com", "abc.com", "localhost"},
			},
			want: true,
		},
		{
			name: "sub domain",
			args: args{
				urlStr:           "https://www.google.com",
				blackListDomains: []string{"google.com", "abc.com"},
			},
			want: true,
		},
		{
			name: "malformed url",
			args: args{
				urlStr:           "httpdadasdadas",
				blackListDomains: []string{"google.com", "abc.com"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckBlacklistDomain(tt.args.urlStr, tt.args.blackListDomains); got != tt.want {
				t.Errorf("CheckBlacklistDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

