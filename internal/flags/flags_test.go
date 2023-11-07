package flags

import (
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	var tests []struct {
		name string
		want *Flags
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFlags(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
