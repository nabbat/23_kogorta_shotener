package envirements

import (
	"reflect"
	"testing"
)

func TestParseEnv(t *testing.T) {
	var tests []struct {
		name string
		want *EnvConfig
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
