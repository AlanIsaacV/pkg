package wday

import (
	"testing"
	"time"
)

var now = time.Now()

func BenchmarkIsWorkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWorkDayToday("CO")
	}
}

func BenchmarkIsWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsWorkDay("CO", now)
	}
}

func BenchmarkNextWorkDay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NextWorkDay("CO")
	}
}

func BenchmarkIsMarketOpen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMarketOpen(time.Now())
	}
}

func BenchmarkIsMarketOpenNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsMarketOpen(now)
	}
}
