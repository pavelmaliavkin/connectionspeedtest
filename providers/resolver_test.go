package providers

import (
	"reflect"
	"testing"
)

type mockProvider struct {
}

func (m mockProvider) GetConnectionSpeed() (downloadMbps_ float64, uploadMbps_ float64, err_ error) {
	return 0, 0, nil
}

func TestGetProvider(t *testing.T) {
	pr := mockProvider{}
	RegisterProvider(`test-provider`, pr)

	type args struct {
		key ProviderKey
	}
	var tests = []struct {
		name  string
		args  args
		want  IProvider
		want1 bool
	}{
		{
			name: "ok",
			args: args{
				key: "test-provider",
			},
			want:  pr,
			want1: true,
		},
		{
			name: "false",
			args: args{
				key: "not-existing",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetProvider(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvider() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetProvider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRegisterProvider(t *testing.T) {

	type args struct {
		key      ProviderKey
		provider IProvider
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				key:      "pr1",
				provider: mockProvider{},
			},
			wantErr: false,
		},
		{
			name: "duplicate",
			args: args{
				key:      "pr1",
				provider: mockProvider{},
			},
			wantErr: true,
		},
		{
			name: "success 2",
			args: args{
				key:      "pr2",
				provider: mockProvider{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterProvider(tt.args.key, tt.args.provider); (err != nil) != tt.wantErr {
				t.Errorf("RegisterProvider() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
