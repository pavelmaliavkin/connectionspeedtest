package netflix

import (
	"fmt"
	"testing"
)

type mockClient struct {
	downloadSpeed float64
	downloadError error
	uploadSpeed   float64
	uploadError   error
}

func (m mockClient) DownloadTest() (float64, error) {
	return m.downloadSpeed, m.downloadError
}

func (m mockClient) UploadTest() (float64, error) {
	return m.uploadSpeed, m.uploadError
}

func Test_provider_GetConnectionSpeed(t *testing.T) {
	type fields struct {
		client IApiClient
	}
	tests := []struct {
		name              string
		fields            fields
		wantDownloadMbps_ float64
		wantUploadMbps_   float64
		wantErr           bool
	}{
		{
			name: "not initialized",
			fields: fields{
				client: nil,
			},
			wantDownloadMbps_: 0,
			wantUploadMbps_:   0,
			wantErr:           true,
		},
		{
			name: "download error",
			fields: fields{
				client: mockClient{
					downloadSpeed: 0,
					downloadError: fmt.Errorf("download error"),
				},
			},
			wantDownloadMbps_: 0,
			wantUploadMbps_:   0,
			wantErr:           true,
		},
		{
			name: "upload error",
			fields: fields{
				client: mockClient{
					uploadSpeed: 0,
					uploadError: fmt.Errorf("download error"),
				},
			},
			wantDownloadMbps_: 0,
			wantUploadMbps_:   0,
			wantErr:           true,
		},
		{
			name: "success",
			fields: fields{
				client: mockClient{
					downloadSpeed: 11,
					uploadSpeed:   12.34,
				},
			},
			wantDownloadMbps_: 11,
			wantUploadMbps_:   12.34,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &provider{
				client: tt.fields.client,
			}
			gotDownloadMbps_, gotUploadMbps_, err := a.GetConnectionSpeed()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConnectionSpeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDownloadMbps_ != tt.wantDownloadMbps_ {
				t.Errorf("GetConnectionSpeed() gotDownloadMbps_ = %v, want %v", gotDownloadMbps_, tt.wantDownloadMbps_)
			}
			if gotUploadMbps_ != tt.wantUploadMbps_ {
				t.Errorf("GetConnectionSpeed() gotUploadMbps_ = %v, want %v", gotUploadMbps_, tt.wantUploadMbps_)
			}
		})
	}
}

func BenchmarkProvider_GetConnectionSpeed(b *testing.B) {
	a := &provider{
		client: getClient(),
	}
	b.ResetTimer()

	a.GetConnectionSpeed()
}
