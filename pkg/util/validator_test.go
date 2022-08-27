package util

import (
	"reflect"
	"testing"
)

package util

import (
"github.com/stretchr/testify/assert"
"reflect"
"testing"
)

func TestIsPhone(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"09foo12322", false},
		{"ab12345678", false},
		{"foo", false},
		{"0123", false},
		{"84912000000", true},
		{"84345333444", true},
		{"842145333444", true},
	}
	for _, test := range tests {
		actual := IsPhone(test.param)
		assert.Equal(t, test.expected, actual, "expect %v, got %v", test.expected, actual)
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.coffee..coffee", false},
		{"foo@bar.bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		actual := IsEmail(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsEmail(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsMobilePhone(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"090988888888", false}, // (09) +
		{"09888888889", false},  // length 11
		{"0123456789", false},
		{"09foo12322", false},
		{"ab12345678", false},
		{"foo", false},
		{"0123", false},
		{"84912000000", true},
		{"84345333444", true},
		{"84888888888", true},
		{"84555577777", true},
		{"84777777777", true},
		{"84277777777", false},
		{"8477777777", false},
		{"0777777777", false},
	}
	for _, test := range tests {
		actual := IsMobilePhone(test.param)
		assert.Equal(t, test.expected, actual)
	}
}

func TestIsValidUsername(t *testing.T) {
	//t.Parallel()
	//fmt.Println(rxUsername.MatchString("foo.bar"))
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo.bar", true},
		{"_", false},
		{"foo_bar", false},
		{"foobar123", true},
		{"123foobar", true},
		{"foobar", true},
		{"09888888888", true},
	}
	for _, test := range tests {
		actual := IsValidUsername(test.param)
		assert.Equal(t, test.expected, actual)
	}
}

func TestValidateURL(t *testing.T) {
	type testCaseURL struct {
		name    string
		testURL string
		want    bool
	}
	tests := []testCaseURL{
		{
			name:    "missing schema",
			testURL: "google.com",
			want:    false,
		},
		{
			name:    "missing host",
			testURL: "https://.com",
			want:    false,
		},
		{
			name:    "missing domain",
			testURL: "https://",
			want:    false,
		},
		{
			name:    "http schema",
			testURL: "http://google.com",
			want:    false,
		},
		{
			name:    "invalid domain",
			testURL: "https://googlecom",
			want:    false,
		},
		{
			name:    "subdomain is localhost",
			testURL: "https://localhost.ghtk.com/img/apple.png",
			want:    true,
		},
		{
			name:    "localhost is a part of domain",
			testURL: "https://localhost.vn/img/apple.png",
			want:    true,
		},
		{
			name:    "localhost url",
			testURL: "https://localhost:8080/img/apple.png",
			want:    false,
		},
		{
			name:    "localhost with subdomain",
			testURL: "https://abc.cdn.localhost:8080/img/apple.png",
			want:    false,
		},
		{
			name:    "valid url",
			testURL: "https://google.com",
			want:    true,
		},
		{
			name:    "valid url with many subdomain",
			testURL: "https://cdn.image.google.com",
			want:    true,
		},
		{
			name:    "valid url with path",
			testURL: "https://cdn.google.com/img/apple.png",
			want:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(subT *testing.T) {
			got := ValidateURL(test.testURL)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expect mismatch, got: %v want: %v", got, test.want)
			}
		})
	}
}

func TestValidateRawURL(t *testing.T) {
	tests := []struct {
		name       string
		testURL    string
		requireSSL bool
		want       bool
	}{
		{
			name:    "missing schema",
			testURL: "google.com",
			want:    false,
		},
		{
			name:    "missing host",
			testURL: "https://.com",
			want:    false,
		},
		{
			name:    "missing domain",
			testURL: "https://",
			want:    false,
		},
		{
			name:       "http schema when require ssl",
			testURL:    "http://google.com",
			requireSSL: true,
			want:       false,
		},
		{
			name:    "http schema when NOT require ssl",
			testURL: "http://google.com",
			want:    true,
		},
		{
			name:    "subdomain is localhost",
			testURL: "https://localhost.ghtk.com/img/apple.png",
			want:    true,
		},
		{
			name:       "localhost url when require SSL",
			testURL:    "https://localhost:8080/img/apple.png",
			requireSSL: true,
			want:       false,
		},
		{
			name:    "localhost url when NOT require SSL",
			testURL: "https://localhost:8080/img/apple.png",
			want:    true,
		},
		{
			name:       "localhost with subdomain",
			testURL:    "https://cdn.localhost:8080/img/apple.png",
			requireSSL: true,
			want:       false,
		},
		{
			name:    "valid url",
			testURL: "https://google.com",
			want:    true,
		},
		{
			name:    "valid url with many subdomain",
			testURL: "https://cdn.image.google.com",
			want:    true,
		},
		{
			name:    "valid url with path",
			testURL: "https://cdn.google.com/img/apple.png",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateRawURL(tt.testURL, tt.requireSSL); got != tt.want {
				t.Errorf("ValidateRawURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
