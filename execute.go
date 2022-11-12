package connectionspeedtest

import (
	"fmt"
	"github.com/pavelmaliavkin/connectionspeedtest/providers"
)

func GetConnectionSpeed(providerKey providers.ProviderKey) (downloadMbps_ float64, uploadMbps_ float64, err_ error) {
	provider, ok := providers.GetProvider(providerKey)
	if !ok {
		return 0, 0, fmt.Errorf("provider with key %s is not supported", providerKey)
	}

	return provider.GetConnectionSpeed()
}
