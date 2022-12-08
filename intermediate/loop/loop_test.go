package loop

import "testing"

func BenchmarkCreateSource(b *testing.B) {
	src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kLoop(src, 2)
	}
}
