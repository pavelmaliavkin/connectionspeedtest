package connectionspeedtest

import (
	"github.com/pavelmaliavkin/connectionspeedtest/providers"
	"testing"
)

type mockProvider struct {
}

func (m mockProvider) GetConnectionSpeed() (downloadMbps_ float64, uploadMbps_ float64, err_ error) {
	return 12, 15, nil
}

func TestGetConnectionSpeed(t *testing.T) {
	providers.RegisterProvider(`test-provider`, mockProvider{})

	type args struct {
		providerKey providers.ProviderKey
	}
	tests := []struct {
		name              string
		args              args
		wantDownloadMbps_ float64
		wantUploadMbps_   float64
		wantErr           bool
	}{
		{
			name: "not registered provider",
			args: args{
				providerKey: "unknown",
			},
			wantDownloadMbps_: 0,
			wantUploadMbps_:   0,
			wantErr:           true,
		},
		{
			name: "success",
			args: args{
				providerKey: "test-provider",
			},
			wantDownloadMbps_: 12,
			wantUploadMbps_:   15,
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDownloadMbps_, gotUploadMbps_, err := GetConnectionSpeed(tt.args.providerKey)
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
