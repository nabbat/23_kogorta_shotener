package shotenermaker

import "testing"

func TestGenerateID(t *testing.T) {
	type args struct {
		fullURL []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateID(tt.args.fullURL); got != tt.want {
				t.Errorf("GenerateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
