package config

import (
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"reflect"
	"testing"
)

func TestSetEnv(t *testing.T) {
	tests := []struct {
		name string
		want *envirements.EnvConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
