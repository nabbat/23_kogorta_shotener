package handlers

import (
	"github.com/gorilla/mux"
	"github.com/nabbat/23_kogorta_shotener/internal/envirements"
	"github.com/nabbat/23_kogorta_shotener/internal/liblog"
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
		log     liblog.Logger
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
			if got := rh.HandleRedirect(tt.args.storage, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequestLoggingMiddleware(t *testing.T) {
	type args struct {
		log liblog.Logger
	}
	tests := []struct {
		name string
		args args
		want mux.MiddlewareFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RequestLoggingMiddleware(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestLoggingMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseLoggingMiddleware(t *testing.T) {
	type args struct {
		log liblog.Logger
	}
	tests := []struct {
		name string
		args args
		want mux.MiddlewareFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResponseLoggingMiddleware(tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseLoggingMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenURLHandler_HandleShortenURL(t *testing.T) {
	type args struct {
		storage *urlstorage.URLStorage
		c       *envirements.EnvConfig
		log     liblog.Logger
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
			if got := sh.HandleShortenURL(tt.args.storage, tt.args.c, tt.args.log); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleShortenURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseWriterWrapper_Size(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		status         int
		size           int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &responseWriterWrapper{
				ResponseWriter: tt.fields.ResponseWriter,
				status:         tt.fields.status,
				size:           tt.fields.size,
			}
			if got := rw.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseWriterWrapper_Status(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		status         int
		size           int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &responseWriterWrapper{
				ResponseWriter: tt.fields.ResponseWriter,
				status:         tt.fields.status,
				size:           tt.fields.size,
			}
			if got := rw.Status(); got != tt.want {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseWriterWrapper_Write(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		status         int
		size           int
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := &responseWriterWrapper{
				ResponseWriter: tt.fields.ResponseWriter,
				status:         tt.fields.status,
				size:           tt.fields.size,
			}
			got, err := rw.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_responseWriterWrapper_WriteHeader(t *testing.T) {
	type fields struct {
		ResponseWriter http.ResponseWriter
		status         int
		size           int
	}
	type args struct {
		status int
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
			rw := &responseWriterWrapper{
				ResponseWriter: tt.fields.ResponseWriter,
				status:         tt.fields.status,
				size:           tt.fields.size,
			}
			rw.WriteHeader(tt.args.status)
		})
	}
}
