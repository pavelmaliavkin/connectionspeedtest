package connectionspeedtest

import (
	"github.com/pavelmaliavkin/connectionspeedtest/providers/netflix"
	"github.com/pavelmaliavkin/connectionspeedtest/providers/ookla"
	"testing"
)

func BenchmarkProvider_Ookla_GetConnectionSpeed(b *testing.B) {
	GetConnectionSpeed(ookla.ProviderKey)
}

func BenchmarkProvider_Netflix_GetConnectionSpeed(b *testing.B) {
	GetConnectionSpeed(netflix.ProviderKey)
}
