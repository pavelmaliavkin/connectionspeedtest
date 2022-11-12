package ookla

import (
	"fmt"
	"github.com/pavelmaliavkin/connectionspeedtest/providers"
)

const ProviderKey providers.ProviderKey = `ookla`

func init() {
	providers.RegisterProvider(ProviderKey, &provider{
		client: getClient(),
	})
}

type provider struct {
	client IApiClient
}

func (a *provider) GetConnectionSpeed() (downloadMbps_ float64, uploadMbps_ float64, err_ error) {
	if a.client == nil {
		return 0, 0, fmt.Errorf(`provider was not initialized properly`)
	}

	downloadMbps_, err_ = a.client.DownloadTest()
	if err_ != nil {
		return 0, 0, fmt.Errorf(`download test failed, err: %w`, err_)
	}

	uploadMbps_, err_ = a.client.UploadTest()
	if err_ != nil {
		return 0, 0, fmt.Errorf(`upload test failed, err: %w`, err_)
	}

	return
}
