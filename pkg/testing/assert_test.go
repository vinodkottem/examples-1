package max

import (
	"testing"

	"github.com/efficientgo/core/testutil"
)

func BenchmarkAssert(b *testing.B) {
	b.ReportAllocs()

	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testutil.Ok(b, err)
	}
}
