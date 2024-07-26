package wday

import (
	"testing"
)

func BenchmarkIsWorkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWorkDayToday("CO")
	}
}

func BenchmarkNextWorkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextWorkDay("CO")
	}
}
