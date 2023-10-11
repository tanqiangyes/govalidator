package govalidator

// Benchmark testing is produced with randomly filled array of 1 million elements

import (
	"math/rand"
	"testing"
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomArray(n int) (res []int) {
	res = make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = randomInt(-1000, 1000)
	}

	return
}

func BenchmarkEach(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn Iterator[int] = func(value int, index int) {
			value = value + index
		}
		Each(data, fn)
	}
}

func BenchmarkReduce(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ReduceIterator[int] = func(init int, val int) int {
			return init + val
		}
		Reduce(data, fn, 0)
	}
}

func BenchmarkMap(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ResultIterator[int] = func(value int, index int) int {
			return value * 3
		}
		_ = Map(data, fn)
	}
}

func BenchmarkFind(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		findElement := 96
		var fn1 ConditionIterator[int] = func(value int, index int) bool {
			return value == findElement
		}
		_ = Find(data, fn1)
	}
}

func BenchmarkFilter(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ConditionIterator[int] = func(value int, index int) bool {
			return value%2 == 0
		}
		_ = Filter(data, fn)
	}
}

func BenchmarkCount(b *testing.B) {
	data := randomArray(1000000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var fn ConditionIterator[int] = func(value int, index int) bool {
			return value%2 == 0
		}
		_ = Count(data, fn)
	}
}
