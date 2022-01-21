// Package hash contains a hashing function for generating deterministic shortcodes.
package hash

import (
	"reflect"
	"testing"
)

func Test_getHash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want [16]byte
	}{
		{
			name: "Valid text",
			args: args{
				text: "http://www.google.com",
			},
			want: [16]byte{237, 100, 106, 51, 52, 202, 137, 31, 211, 70, 125, 177, 49, 55, 33, 64},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHash(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GenerateShortCode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid text",
			args: args{
				text: "http://www.google.com",
			},
			want: "ed646a33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShortCode(tt.args.text); got != tt.want {
				t.Errorf("generateShortCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
