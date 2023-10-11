package storage

import (
	"reflect"
	"testing"
)

func TestNewURLStorage(t *testing.T) {
	tests := []struct {
		name string
		want *URLStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewURLStorage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLStorage_AddURL(t *testing.T) {
	type fields struct {
		urlMap map[string]string
	}
	type args struct {
		shortURL    string
		originalURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &URLStorage{
				urlMap: tt.fields.urlMap,
			}
			storage.AddURL(tt.args.shortURL, tt.args.originalURL)
		})
	}
}

func TestURLStorage_GetOriginalURL(t *testing.T) {
	type fields struct {
		urlMap map[string]string
	}
	type args struct {
		shortURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &URLStorage{
				urlMap: tt.fields.urlMap,
			}
			if got := storage.GetOriginalURL(tt.args.shortURL); got != tt.want {
				t.Errorf("GetOriginalURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
