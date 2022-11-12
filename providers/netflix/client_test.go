package netflix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_DownloadTest(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &client{}
			got, err := a.DownloadTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEqual(t, got, 0)
		})
	}
}

func Test_client_UploadTest(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &client{}
			got, err := a.UploadTest()
			if (err != nil) != tt.wantErr {
				t.Errorf("UploadTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEqual(t, got, 0)
		})
	}
}

func Test_getClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getClient()
			assert.NotNil(t, got)
		})
	}
}
