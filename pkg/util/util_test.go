package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildServerAddress(t *testing.T) {
	type args struct {
		host string
		port int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "fullHostAndPort", args: args{host: "127.0.0.1", port: 8080}, want: "127.0.0.1:8080"},
		{name: "onlyPort", args: args{host: "", port: 8080}, want: ":8080"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetServerAddress(tt.args.host, tt.args.port); got != tt.want {
				t.Errorf("GetServerAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCensorString() {
	fmt.Println(CensorString("this_is_a_secret_text"))
	// Output: this_is_***
}

func TestCensorString(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "normal text", arg: "this_is_secret", want: "this_is***"},
		{name: "short text", arg: "1", want: "***"},
		{name: "long text", arg: "this_is_a_very_very_very_long_text", want: "this_is_***"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CensorString(tt.arg); got != tt.want {
				t.Errorf("CensorString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCensorPhone() {
	fmt.Println(CensorPhone("0961211789", 2, 3))
	// Output: 09XXXXX789
}

func TestCensorPhone(t *testing.T) {
	type args struct {
		phone  string
		prefix int
		suffix int
	}
	tests := []struct {
		name string
		arg  args
		want string
	}{
		{name: "normal phone", arg: args{"0961211789", 2, 3}, want: "09XXXXX789"},
		{name: "suffix is more than phone length", arg: args{"0961211", 1, 12}, want: "0961211"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CensorPhone(tt.arg.phone, tt.arg.prefix, tt.arg.suffix); got != tt.want {
				t.Errorf("CensorPhone () = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMySQLDataSourceName(t *testing.T) {
	type args struct {
		host     string
		port     int
		protocol string
		username string
		password string
		dbname   string
		params   string
	}

	test := args{
		host:     "10.10.110.69",
		port:     7890,
		protocol: "tcp",
		username: "admin",
		password: "pwd_12345",
		dbname:   "gmicro",
		params:   "check=true",
	}

	want := "admin:pwd_12345@tcp(10.10.110.69:7890)/gmicro?check=true"

	if got := GetMySQLDataSourceName(test.host, test.port, test.protocol,
		test.username, test.password, test.dbname, test.params); got != want {
		t.Errorf("GetMySQLDataSourceName () = %v, want %v", got, want)
	}
}

func TestIsNil(t *testing.T) {
	type args struct {
		i interface{}
	}

	var nilSlice []int
	var nilMap map[string]string
	var nilPtr *int
	var nilChan chan int
	var nilInterface interface{}
	var nilFunc func()
	var nilPtrStruct *struct{}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{},
			want: true,
		},
		{
			name: "slice",
			args: args{i: make([]int, 3)},
			want: false,
		},
		{
			name: "slice nil",
			args: args{i: nilSlice},
			want: true,
		},
		{
			name: "map",
			args: args{i: map[string]int{"a": 1, "b": 2}},
			want: false,
		},
		{
			name: "map nil",
			args: args{i: nilMap},
			want: true,
		},
		{
			name: "interface nil",
			args: args{i: nilInterface},
			want: true,
		},
		{
			name: "func nil",
			args: args{i: nilFunc},
			want: true,
		},
		{
			name: "pointer struct nil",
			args: args{i: nilPtrStruct},
			want: true,
		},
		{
			name: "array",
			args: args{i: []int{2, 3, 4, 5}},
			want: false,
		},
		{
			name: "pointer nil",
			args: args{i: nilPtr},
			want: true,
		},
		{
			name: "channel",
			args: args{i: make(chan int)},
			want: false,
		},
		{
			name: "channel nil",
			args: args{i: nilChan},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.i); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsStringSliceContains(t *testing.T) {
	arrStr := []string{"Hello", "world", "OK"}
	findStr := "world"
	rs := IsStringSliceContains(arrStr, findStr)
	assert.True(t, rs)

	findStr = "KO"
	rsF := IsStringSliceContains(arrStr, findStr)
	assert.False(t, rsF)
}
