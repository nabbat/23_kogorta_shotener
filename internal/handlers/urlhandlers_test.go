package handlers

import (
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	urlstorage "github.com/nabbat/23_kogorta_shotener/internal/storage"
	"net/http"
	"reflect"
	"testing"
)

func TestPanicHandler(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PanicHandler(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PanicHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedirectHandler_HandleRedirect(t *testing.T) {
	type args struct {
		storage *urlstorage.URLStorage
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rh := &RedirectHandler{}
			if got := rh.HandleRedirect(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenURLHandler_HandleShortenURL(t *testing.T) {
	type args struct {
		storage *urlstorage.URLStorage
		c       *envirements.EnvConfig
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &ShortenURLHandler{}
			if got := sh.HandleShortenURL(tt.args.storage, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleShortenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
