package govalidator

import (
	"testing"
)

func BenchmarkToBoolean(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ToBoolean("False    ")
	}
}

func BenchmarkToInt(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ToInt("-22342342.2342")
	}
}

func BenchmarkToNumber(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ToNumber[string, float64]("-22342342.2342")
	}
}

func BenchmarkToString(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		ToString(randomArray(1000000))
	}
}

func BenchmarkToJson(b *testing.B) {
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, _ = ToJSON(randomArray(1000000))
	}
}
