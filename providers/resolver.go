package providers

import "fmt"

type ProviderKey = string

var providers map[ProviderKey]IProvider

func init() {
	providers = make(map[ProviderKey]IProvider, 0)
}

type IProvider interface {
	GetConnectionSpeed() (downloadMbps_ float64, uploadMbps_ float64, err_ error)
}

func RegisterProvider(key ProviderKey, provider IProvider) error {
	_, ok := providers[key]
	if ok {
		return fmt.Errorf("provider with key %s already registered", key)
	}

	providers[key] = provider
	return nil
}

func GetProvider(key ProviderKey) (IProvider, bool) {
	provider, ok := providers[key]
	return provider, ok
}
