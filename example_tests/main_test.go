package main

import (
	"testing"
)

/*
func Test_Calc(t *testing.T) {
	t.Parallel()

	t.Run("Sum is calculated cool!", func(t *testing.T) {
		t.Parallel()
		result := 3
		res := CalcSum(1, 2)
		assert.Equal(t, result, res)
	})
}

*/

func BenchmarkAtomicCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AtomicCounter()
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MutexCounter()
	}
}
